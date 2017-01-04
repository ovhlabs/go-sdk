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

// Informations related to a line offer
type TelephonyLineOffer struct {

	// The offer description
	Description string `json:"description,omitempty"`

	// The offer name
	Name string `json:"name,omitempty"`

	Price OrderPrice `json:"price,omitempty"`
}