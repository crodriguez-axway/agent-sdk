package agent

import (
	"github.com/Axway/agent-sdk/pkg/apic/apiserver/models/management/v1alpha1"
	"github.com/Axway/agent-sdk/pkg/config"
	"github.com/Axway/agent-sdk/pkg/util/log"
)

func createTraceabilityAgentResource(config v1alpha1.TraceabilityAgentSpecConfig, logging v1alpha1.DiscoveryAgentSpecLogging, gatewayType string) {
	// The traceability agent resource needs to be created
	agentResource := v1alpha1.TraceabilityAgent{}

	agentResource.Spec.Config = config
	agentResource.Spec.Logging = logging
	agentResource.Spec.DataplaneType = gatewayType
	agentResource.Name = agent.cfg.GetAgentName()

	log.Debug("Creating the traceability agent resource")
	createAgentResource(&agentResource)

	log.Debug("Updating the traceability agent status sub-resource")
	updateAgentStatusAPI(&agentResource, v1alpha1.TraceabilityAgentResource)
}

func createTraceabilityAgentStatusResource(status, message string) *v1alpha1.TraceabilityAgent {
	agentRes := v1alpha1.TraceabilityAgent{}
	agentRes.Name = agent.cfg.GetAgentName()
	agentRes.Status.Version = config.AgentVersion
	agentRes.Status.State = status
	agentRes.Status.Message = message

	return &agentRes
}

func mergeTraceabilityAgentWithConfig(cfg *config.CentralConfiguration) {
	// Nothing to merge
}