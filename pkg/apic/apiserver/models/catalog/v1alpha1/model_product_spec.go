/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha1

// ProductSpec struct for ProductSpec
type ProductSpec struct {
	// description of the Product.
	Description string           `json:"description,omitempty"`
	Asset       ProductSpecAsset `json:"asset,omitempty"`
}