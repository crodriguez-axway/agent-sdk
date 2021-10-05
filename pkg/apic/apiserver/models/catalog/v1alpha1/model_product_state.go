/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha1

// ProductState the model 'ProductState'
type ProductState string

// List of ProductState
const (
	DRAFT      ProductState = "draft"
	ACTIVE     ProductState = "active"
	DEPRECATED ProductState = "deprecated"
	ARCHIVED   ProductState = "archived"
)
