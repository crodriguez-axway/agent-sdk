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

// DocumentationStatus struct for DocumentationStatus
type DocumentationStatus struct {
	Type  string `json:"type,omitempty"`
	Level string `json:"level,omitempty"`
	// A brief CamelCase message indicating details of the unexpected error while processing the AssetRelease update.
	Title string `json:"title,omitempty"`
	// Descriptive details of the error.
	Detail string `json:"detail,omitempty"`
	// Time when the update occurred.
	Timestamp time.Time `json:"timestamp,omitempty"`
	// Custom reason details.
	Meta map[string]map[string]interface{} `json:"meta,omitempty"`
}
