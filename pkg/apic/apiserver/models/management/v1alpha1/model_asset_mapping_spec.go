/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha1

// AssetMappingSpec struct for AssetMappingSpec
type AssetMappingSpec struct {
	// Reference to the executed AssetMappingTemplate.
	AssetMappingTemplate string                 `json:"assetMappingTemplate"`
	Inputs               AssetMappingSpecInputs `json:"inputs"`
}