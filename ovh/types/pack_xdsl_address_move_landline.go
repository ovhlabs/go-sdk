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

// PackXdslAddressMoveLandline The parameters needed to activate the access on a landline
type PackXdslAddressMoveLandline struct {

	// LineNumber The number of the landline
	LineNumber string `json:"lineNumber,omitempty"`

	// PortLineNumber Whether or not tha lineNumber should be ported to OVH, if eligibile
	PortLineNumber bool `json:"portLineNumber,omitempty"`

	// Rio A token to prove the ownership of the line number, needed to port the number
	Rio string `json:"rio,omitempty"`

	// Status The status of the landline
	Status string `json:"status,omitempty"`

	// Unbundling The unbundling of the landline
	Unbundling string `json:"unbundling,omitempty"`
}
