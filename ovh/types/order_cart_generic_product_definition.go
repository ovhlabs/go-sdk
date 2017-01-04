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

// Representation of a generic product
type OrderCartGenericProductDefinition struct {

	// Product offer identifier
	PlanCode string `json:"planCode,omitempty"`

	// Prices of the product offer
	Prices []OrderCartGenericProductPricing `json:"prices,omitempty"`

	// Name of the product
	ProductName string `json:"productName,omitempty"`

	// Product type
	ProductType string `json:"productType,omitempty"`
}