package traceability

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"git.ecd.axway.org/apigov/apic_agents_sdk/pkg/agent"
	"git.ecd.axway.org/apigov/apic_agents_sdk/pkg/config"
	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/outputs"
	"github.com/elastic/beats/v7/libbeat/outputs/outil"
	"github.com/elastic/beats/v7/libbeat/outputs/transport"
	"github.com/elastic/beats/v7/libbeat/publisher"
)

// HTTPClient struct
type HTTPClient struct {
	Connection
	tlsConfig        *transport.TLSConfig
	compressionLevel int
	proxyURL         *url.URL
	observer         outputs.Observer
	headers          map[string]string
	beatInfo         beat.Info
}

// HTTPClientSettings struct
type HTTPClientSettings struct {
	BeatInfo         beat.Info
	URL              string
	Proxy            *url.URL
	TLS              *transport.TLSConfig
	Index            outil.Selector
	Pipeline         *outil.Selector
	Timeout          time.Duration
	CompressionLevel int
	Observer         outputs.Observer
	Headers          map[string]string
}

// Connection struct
type Connection struct {
	URL       string
	http      *http.Client
	connected bool
	encoder   bodyEncoder
}

type eventRaw map[string]json.RawMessage

type event struct {
	Timestamp time.Time         `json:"@timestamp"`
	Metadata  httpEventMetadata `json:"@metadata"`
	Fields    common.MapStr     `json:"-"`
}

// Meta defines common event metadata to be stored in '@metadata'
type httpEventMetadata struct {
	Beat    string `json:"beat"`
	Type    string `json:"type"`
	Version string `json:"version"`
}

// NewHTTPClient instantiate a client.
func NewHTTPClient(s HTTPClientSettings) (*HTTPClient, error) {
	proxy := http.ProxyFromEnvironment
	if s.Proxy != nil {
		proxy = http.ProxyURL(s.Proxy)
	}

	var dialer, tlsDialer transport.Dialer
	var err error

	dialer = transport.NetDialer(s.Timeout)
	tlsDialer, err = transport.TLSDialer(dialer, s.TLS, s.Timeout)
	if err != nil {
		return nil, err
	}

	if st := s.Observer; st != nil {
		dialer = transport.StatsDialer(dialer, st)
		tlsDialer = transport.StatsDialer(tlsDialer, st)
	}

	var encoder bodyEncoder
	compression := s.CompressionLevel
	if compression == 0 {
		encoder = newJSONEncoder(nil)
	} else {
		encoder, err = newGzipEncoder(compression, nil)
		if err != nil {
			return nil, err
		}
	}
	client := &HTTPClient{
		Connection: Connection{
			URL: s.URL,
			http: &http.Client{
				Transport: &http.Transport{
					Dial:    dialer.Dial,
					DialTLS: tlsDialer.Dial,
					Proxy:   proxy,
				},
				Timeout: s.Timeout,
			},
			encoder: encoder,
		},
		compressionLevel: compression,
		proxyURL:         s.Proxy,
		headers:          s.Headers,
		beatInfo:         s.BeatInfo,
	}

	return client, nil
}

// Connect establishes a connection to the clients sink.
func (client *HTTPClient) Connect() error {
	client.Connection.connected = true
	return nil
}

// Close publish a single event to output.
func (client *HTTPClient) Close() error {
	client.Connection.connected = false
	return nil
}

// Publish sends events to the clients sink.
func (client *HTTPClient) Publish(batch publisher.Batch) error {
	events := batch.Events()
	rest, err := client.publishEvents(events)
	if len(rest) == 0 {
		batch.ACK()
	} else {
		batch.RetryEvents(rest)
	}
	return err
}

func (client *HTTPClient) String() string {
	return client.URL
}

// Clone clones a client.
func (client *HTTPClient) Clone() *HTTPClient {
	c, _ := NewHTTPClient(
		HTTPClientSettings{
			BeatInfo:         client.beatInfo,
			URL:              client.URL,
			Proxy:            client.proxyURL,
			TLS:              client.tlsConfig,
			Timeout:          client.http.Timeout,
			CompressionLevel: client.compressionLevel,
			Headers:          client.headers,
		},
	)
	return c
}

// publishEvents - posts all events to the http endpoint.
func (client *HTTPClient) publishEvents(data []publisher.Event) ([]publisher.Event, error) {
	if len(data) == 0 {
		return nil, nil
	}

	if !client.connected {
		return data, ErrHTTPNotConnected
	}

	var events = make([]eventRaw, len(data))
	for i, event := range data {
		events[i] = client.makeHTTPEvent(&event.Content)
	}
	status, _, err := client.request(events, client.headers)
	if err != nil {
		if err == ErrJSONEncodeFailed {
			debugf("Failed to publish event: %s", err.Error())
			return nil, nil
		}
	}
	switch {
	case status == 500 || status == 400: //server error or bad input, don't retry
		debugf("Failed to publish event: received status code %d", status)
		return nil, nil
	case status >= 300:
		// retry
		return data, err
	}

	return nil, nil
}

func (conn *Connection) request(body interface{}, headers map[string]string) (int, []byte, error) {
	urlStr := conn.URL
	if strings.HasSuffix(urlStr, "/") {
		urlStr = strings.TrimSuffix(urlStr, "/")
	}

	if err := conn.encoder.Marshal(body); err != nil {
		return 0, nil, ErrJSONEncodeFailed
	}
	return conn.execRequest(urlStr, conn.encoder.Reader(), headers)
}

func (conn *Connection) execRequest(url string, body io.Reader, headers map[string]string) (int, []byte, error) {
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return 0, nil, err
	}

	err = conn.addHeaders(&req.Header, body)
	if err != nil {
		return 0, nil, err
	}

	return conn.execHTTPRequest(req, headers)
}

func (conn *Connection) addHeaders(header *http.Header, body io.Reader) error {
	token, err := agent.GetCentralAuthToken()
	if err != nil {
		return err
	}

	header.Add("Authorization", "Bearer "+token)
	header.Add("axway-target-flow", "api-central-v8")
	header.Add("User-Agent", config.AgentTypeName+"/"+config.AgentVersion)
	if body != nil {
		conn.encoder.AddHeader(header)
	}
	return nil
}

func (conn *Connection) execHTTPRequest(req *http.Request, headers map[string]string) (int, []byte, error) {
	for key, value := range headers {
		req.Header.Add(key, value)
	}

	resp, err := conn.http.Do(req)
	if err != nil {
		conn.connected = false
		return 0, nil, err
	}
	defer closing(resp.Body)

	status := resp.StatusCode
	if status >= 300 {
		conn.connected = false
		return status, nil, fmt.Errorf("%v", resp.Status)
	}
	obj, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		conn.connected = false
		return status, nil, err
	}
	return status, obj, nil
}

func closing(c io.Closer) {
	c.Close()
}

func (client *HTTPClient) makeHTTPEvent(v *beat.Event) map[string]json.RawMessage {
	type event0 event
	e := event{
		Timestamp: v.Timestamp.UTC(),
		Metadata: httpEventMetadata{
			Beat:    client.beatInfo.Beat + "_http",
			Version: client.beatInfo.Version,
			Type:    "_doc",
		},
		Fields: v.Fields,
	}
	b, _ := json.Marshal(event0(e))

	var eventMap map[string]json.RawMessage
	json.Unmarshal(b, &eventMap)
	for j, k := range e.Fields {
		b, _ = json.Marshal(k)
		eventMap[j] = b
	}
	return eventMap
}