/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha1

// ResourceSpec struct for ResourceSpec
type ResourceSpec struct {
	Type string `json:"type"`
	// Resource content.
	Content string `json:"content"`
}