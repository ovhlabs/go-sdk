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

// EmailExchangeOutlookVersions Availability of outlook version
type EmailExchangeOutlookVersions struct {
	OutlookLanguage string `json:"outlookLanguage,omitempty"`

	OutlookVersion string `json:"outlookVersion,omitempty"`

	Status bool `json:"status,omitempty"`
}
