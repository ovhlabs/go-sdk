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

// Stock
type CloudReservedStock struct {

	// Error message, if any
	Error_ string `json:"error,omitempty"`

	// Flavor id
	FlavorId string `json:"flavorId,omitempty"`

	// Flavor name
	FlavorName string `json:"flavorName,omitempty"`

	// Number of available instances
	Quantity int64 `json:"quantity,omitempty"`
}