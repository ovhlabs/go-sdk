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

// OrderTelephonyAccessoriesPost ...
type OrderTelephonyAccessoriesPost struct {
	Accessories []string `json:"accessories,omitempty"`

	MondialRelayID string `json:"mondialRelayId,omitempty"`

	Retractation bool `json:"retractation,omitempty"`

	ShippingContactID int64 `json:"shippingContactId,omitempty"`
}
