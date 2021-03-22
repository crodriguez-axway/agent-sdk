/*
 * Amplify Unified Catalog APIs
 *
 * APIs for Amplify Unified Catalog
 *
 * API version: 1.43.0
 * Contact: support@axway.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package unifiedcatalog
// CatalogItemRevision struct for CatalogItemRevision
type CatalogItemRevision struct {
	// Generated identifier for the resource
	Id string `json:"id,omitempty"`
	Metadata AuditMetadata `json:"metadata,omitempty"`
	Number int32 `json:"number,omitempty"`
	// Version name for the revision.
	Version string `json:"version"`
	State string `json:"state"`
	Status string `json:"status,omitempty"`
}