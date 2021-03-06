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

// CloudSSHKey SshKey
type CloudSSHKey struct {

	// ID SSH key id
	ID string `json:"id,omitempty"`

	// Name SSH key name
	Name string `json:"name,omitempty"`

	// PublicKey SSH public key
	PublicKey string `json:"publicKey,omitempty"`

	// Regions SSH key regions
	Regions []string `json:"regions,omitempty"`
}
