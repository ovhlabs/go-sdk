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

// OrderCatalogPricingDefault Describe default pricings
type OrderCatalogPricingDefault struct {

	// TDefault Information about default pricing
	TDefault []*OrderCatalogPricing `json:"default,omitempty"`
}
