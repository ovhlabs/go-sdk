/*
 * OVH API - EU
 *
 * Build your own OVH world.
 *
 * OpenAPI spec version: 1.0.0
 * Contact: api-subscribe@ml.ovh.net
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package types

// SupplyMondialRelayReturn Status and Mondial Relay Point Details
type SupplyMondialRelayReturn struct {

	// TError Error
	TError string `json:"error,omitempty"`

	Result *SupplyMondialRelayResult `json:"result,omitempty"`

	// Status Request status
	Status string `json:"status,omitempty"`
}
