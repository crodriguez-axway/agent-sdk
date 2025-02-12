package resource

import (
	"strings"

	v1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
	"github.com/Axway/agent-sdk/pkg/apic/apiserver/models/management/v1alpha1"
	"github.com/Axway/agent-sdk/pkg/config"
)

func discoveryAgent(res *v1.ResourceInstance) *v1alpha1.DiscoveryAgent {
	agentRes := &v1alpha1.DiscoveryAgent{}
	agentRes.FromInstance(res)

	return agentRes
}

func createDiscoveryAgentStatusResource(agentName, status, prevStatus, message string) *v1alpha1.DiscoveryAgent {
	agentRes := v1alpha1.DiscoveryAgent{}
	agentRes.Name = agentName
	agentRes.Status.Version = config.AgentVersion
	agentRes.Status.LatestAvailableVersion = config.AgentLatestVersion
	agentRes.Status.State = status
	agentRes.Status.PreviousState = prevStatus
	agentRes.Status.Message = message
	agentRes.Status.LastActivityTime = getTimestamp()
	agentRes.Status.SdkVersion = config.SDKVersion

	return &agentRes
}

func mergeDiscoveryAgentWithConfig(agentRes *v1.ResourceInstance, cfg *config.CentralConfiguration) {
	da := discoveryAgent(agentRes)
	resCfgAdditionalTags := strings.Join(da.Spec.Config.AdditionalTags, ",")
	resCfgTeamName := da.Spec.Config.OwningTeam
	resCfgLogLevel := da.Spec.Logging.Level
	applyResConfigToCentralConfig(cfg, resCfgAdditionalTags, resCfgTeamName, resCfgLogLevel)
}
