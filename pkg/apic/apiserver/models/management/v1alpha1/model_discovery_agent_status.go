/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha1

import (
	// GENERATE: The following code has been modified after code generation
	// 	"time"
	time "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

// DiscoveryAgentStatus struct for DiscoveryAgentStatus
type DiscoveryAgentStatus struct {
	// Version name for the agent revision.
	Version string `json:"version,omitempty"`
	// Latest available version for the agent revision.
	LatestAvailableVersion string `json:"latestAvailableVersion,omitempty"`
	// Agent status:  * running - Passed all health checks.  Up and running  * stopped - Agent is not running  * failed - Failed health checks  * unhealthy - Agent is running with health check failure
	State string `json:"state,omitempty"`
	// A way to communicate details about the current status by the agent
	Message string `json:"message,omitempty"`
	// The last updated event timestamp provided by the agent
	LastActivityTime time.Time `json:"lastActivityTime,omitempty"`
	// Version name for the SDK revision.
	SdkVersion string `json:"sdkVersion,omitempty"`
}
