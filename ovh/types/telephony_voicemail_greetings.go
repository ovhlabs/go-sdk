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

// TelephonyVoicemailGreetings Voicemail greeting
type TelephonyVoicemailGreetings struct {

	// Callee Callee number in international format
	Callee string `json:"callee,omitempty"`

	// Dir Customized greeting voicemail directory
	Dir string `json:"dir,omitempty"`

	// ID Uniq customized greeting identifier
	ID int64 `json:"id,omitempty"`
}
