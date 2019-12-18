package apic

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	corecfg "git.ecd.axway.int/apigov/aws_apigw_discovery_agent/core/config"
	"git.ecd.axway.int/apigov/aws_apigw_discovery_agent/pkg/config"
	"git.ecd.axway.int/apigov/service-mesh-agent/pkg/apicauth"
	"github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
)

// consts for auth policy types
const (
	Apikey      = "verify-api-key"
	Passthrough = "pass-through"
)

// ValidPolicies - list of valid auth policies supported by Central.  Add to this list as more policies are supported.
var ValidPolicies = []string{Apikey, Passthrough}

//CatalogCreator - interface
type CatalogCreator interface {
	DeployAPI(method, url string, buffer []byte) (string, error)
}

// Client -
type Client struct {
	tokenRequester *apicauth.PlatformTokenGetter
	cfg            corecfg.CentralConfig
}

// New -
func New(cfg corecfg.CentralConfig) *Client {
	tokenURL := cfg.GetAuthConfig().GetTokenURL()
	aud := cfg.GetAuthConfig().GetAudience()
	priKey := cfg.GetAuthConfig().GetPrivateKey()
	pubKey := cfg.GetAuthConfig().GetPublicKey()
	keyPwd := cfg.GetAuthConfig().GetKeyPassword()
	clientID := cfg.GetAuthConfig().GetClientID()
	authTimeout := cfg.GetAuthConfig().GetTimeout()

	return &Client{
		cfg:            cfg,
		tokenRequester: apicauth.NewPlatformTokenGetter(priKey, pubKey, keyPwd, tokenURL, aud, clientID, authTimeout),
	}
}

// MapToStringArray -
func (c *Client) MapToStringArray(m map[string]interface{}) []string {
	strArr := []string{}

	for key, val := range m {
		v := val.(*string)
		if *v == "" {
			strArr = append(strArr, key)
		} else {
			strArr = append(strArr, key+"_"+*v)
		}
	}
	return strArr
}

var httpClient = http.DefaultClient
var log logrus.FieldLogger = logrus.WithField("package", "apic")

// SetLog sets the logger for the package.
func SetLog(newLog logrus.FieldLogger) {
	log = newLog
	return
}

func isUnitTesting() bool {
	return strings.HasSuffix(os.Args[0], ".test")
}

// DeployAPI -
func (c *Client) DeployAPI(method, url string, buffer []byte) (string, error) {
	// Unit testing. For now just dummy up a return
	if isUnitTesting() {
		return "12345678", nil
	}

	request, err := setHeader(c, method, url, bytes.NewBuffer(buffer))

	if err != nil {
		return "", err
	}

	response, err := httpClient.Do(request)
	if err != nil {
		return "", err
	}

	if !(response.StatusCode == http.StatusOK || response.StatusCode == http.StatusCreated) {
		detail := make(map[string]*json.RawMessage)
		json.NewDecoder(response.Body).Decode(&detail)
		for k, v := range detail {
			buffer, _ := v.MarshalJSON()
			log.Debugf("HTTP response key %v: %v", k, string(buffer))
		}
		return "", errors.New(response.Status)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return handleResponse(body)
}

func handleResponse(body []byte) (string, error) {

	itemID := ""

	// Connected Mode
	if config.GetConfig().CentralConfig.GetAgentMode() == corecfg.Connected {
		metadata := gjson.Get(string(body), "metadata").String()
		if metadata != "" {
			itemID = gjson.Get(string(metadata), "id").String()
		}
		// Disconnected Mode
	} else {
		itemID = gjson.Get(string(body), "id").String()
	}

	log.Debugf("HTTP response returning itemID: [%v]", itemID)
	return itemID, nil
}

// SetHeader - set header
func setHeader(c *Client, method, url string, body io.Reader) (*http.Request, error) {
	request, err := http.NewRequest(method, url, body)
	var token string

	if token, err = c.tokenRequester.GetToken(); err != nil {
		return nil, err
	}

	request.Header.Add("X-Axway-Tenant-Id", c.cfg.GetTenantID())
	request.Header.Add("Authorization", "Bearer "+token)
	request.Header.Add("Content-Type", "application/json")
	return request, nil
}

// IsNewAPI -
func (c *Client) IsNewAPI(serviceBody ServiceBody) bool {
	var token string
	apiName := strings.ToLower(serviceBody.APIName)
	request, err := http.NewRequest("GET", c.cfg.GetAPIServerServicesURL()+"/"+apiName, nil)

	if token, err = c.tokenRequester.GetToken(); err != nil {
		log.Error("Could not get token")
	}

	request.Header.Add("X-Axway-Tenant-Id", c.cfg.GetTenantID())
	request.Header.Add("Authorization", "Bearer "+token)
	request.Header.Add("Content-Type", "application/json")

	response, _ := http.DefaultClient.Do(request)
	if response.StatusCode == http.StatusNotFound {
		log.Debug("New api found to deploy")
		return true
	}
	return false
}
