/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha1

// AccessRequestSpec struct for AccessRequestSpec
type AccessRequestSpec struct {
	// The name of an APIServiceInstance resource that specifies where the API is deployed.
	ApiServiceInstance string `json:"apiServiceInstance"`
	// The AccessRequest
	AccessRequest string `json:"accessRequest"`
	// The value that matches the AccessRequestDefinition schema linked to the referenced APIServiceInstance.
	Data map[string]interface{} `json:"data"`
}
