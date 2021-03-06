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

// IPMitigationProfile Mitigation profile for your ip
type IPMitigationProfile struct {

	// AutoMitigationTimeOut Delay to wait before remove ip from auto mitigation after an attack
	AutoMitigationTimeOut int64 `json:"autoMitigationTimeOut,omitempty"`

	IPMitigationProfile string `json:"ipMitigationProfile,omitempty"`

	// State Current state of your mitigation profile
	State string `json:"state,omitempty"`
}
