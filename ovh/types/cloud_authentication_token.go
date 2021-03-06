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

// CloudAuthenticationToken Token
type CloudAuthenticationToken struct {
	XAuthToken string `json:"X-Auth-Token,omitempty"`

	Token *CloudAuthenticationOpenstackToken `json:"token,omitempty"`
}
