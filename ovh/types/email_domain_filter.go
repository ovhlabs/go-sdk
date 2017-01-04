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

// Filter List
type EmailDomainFilter struct {

	// Action of filter
	Action string `json:"action,omitempty"`

	// Action parameter of filter
	ActionParam string `json:"actionParam,omitempty"`

	// If true filter is active
	Active bool `json:"active,omitempty"`

	// Domain name of filter
	Domain string `json:"domain,omitempty"`

	// Filter name
	Name string `json:"name,omitempty"`

	// Account name of filter
	Pop string `json:"pop,omitempty"`

	// Priority of filter
	Priority int64 `json:"priority,omitempty"`
}