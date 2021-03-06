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

// MeInstallationTemplatePost ...
type MeInstallationTemplatePost struct {
	BaseTemplateName string `json:"baseTemplateName,omitempty"`

	DefaultLanguage string `json:"defaultLanguage,omitempty"`

	Name string `json:"name,omitempty"`
}
