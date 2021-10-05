package watchmanager

import (
	"context"
	"fmt"

	"github.com/Axway/agent-sdk/pkg/watchmanager/proto"

	"github.com/Axway/agent-sdk/pkg/config"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Manager interface {
	RegisterWatch(watchConfig Config, eventChannel chan *proto.Event) (context.Context, error)
}

type watchManager struct {
	centralConfig config.CentralConfig
	tokenGetter   TokenGetter
	clientMap     map[string]*watchClient
	connection    *grpc.ClientConn
}

func New(centralConfig config.CentralConfig, tokenGetter TokenGetter) Manager {

	manager := &watchManager{
		centralConfig: centralConfig,
		tokenGetter:   tokenGetter,
		clientMap:     make(map[string]*watchClient),
	}
	manager.connection, _ = manager.createConnection()
	return manager
}

func (m *watchManager) RegisterWatch(watchConfig Config, eventChannel chan *proto.Event) (context.Context, error) {
	svcClient := proto.NewWatchServiceClient(m.connection)
	watchRequest := m.createWatchRequest(watchConfig)
	stream, err := svcClient.CreateWatch(context.Background(), watchRequest)
	if err != nil {
		return nil, err
	}

	client := &watchClient{
		config: watchConfig,
		stream: stream,
	}

	uuiduuid, _ := uuid.NewUUID()
	m.clientMap[uuiduuid.String()] = client

	client.processEvents(eventChannel)

	return stream.Context(), nil
}

func (m *watchManager) createConnection() (*grpc.ClientConn, error) {
	address := fmt.Sprintf("%s:%s", "localhost", "8080")

	rpcCredential := newRPCAuth(m.centralConfig.GetTenantID(), m.tokenGetter)
	var dialOptions []grpc.DialOption
	dialOptions = append(dialOptions,
		grpc.WithPerRPCCredentials(rpcCredential),
	)

	if m.centralConfig.GetTLSConfig() != nil {
		tlsConfig := m.centralConfig.GetTLSConfig().BuildTLSConfig()
		dialOptions = append(dialOptions,
			grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)),
		)
	} else {
		dialOptions = append(dialOptions,
			grpc.WithInsecure(),
		)
	}

	return grpc.Dial(address, dialOptions...)
}

func (m *watchManager) createWatchRequest(config Config) *proto.Request {
	triggerEventTypes := make([]proto.Trigger_Type, 0)
	for _, eventType := range config.EventTypes {
		if val, ok := proto.Trigger_Type_value[eventType]; ok {
			triggerEventTypes = append(triggerEventTypes, proto.Trigger_Type(val))
		}
	}

	trigger := &proto.Trigger{
		Group: config.Group,
		Kind:  config.Kind,
		Name:  config.Name,
		Type:  []proto.Trigger_Type{proto.Trigger_CREATED, proto.Trigger_UPDATED, proto.Trigger_DELETED},
	}

	if config.ScopeKind != "" || config.Scope != "" {
		trigger.Scope = &proto.Trigger_Scope{Kind: config.ScopeKind, Name: config.Scope}
	}

	return &proto.Request{
		Triggers: []*proto.Trigger{trigger},
	}
}