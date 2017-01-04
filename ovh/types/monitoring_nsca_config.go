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

// Configuration for the NSCA receiver
type MonitoringNscaConfig struct {

	// true to enable the NSCA receiver (for passive checks)
	Enabled bool `json:"enabled,omitempty"`

	// The kind of encryption (only kind 1 is supported, 0 is for no encryption)
	Encryption int64 `json:"encryption,omitempty"`

	// The encryption key
	Key string `json:"key,omitempty"`
}