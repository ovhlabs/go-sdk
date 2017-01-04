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

// MetricsAccount
type DbaasQueueMetricsAccount struct {

	// OpenTSDB host url
	Host string `json:"host,omitempty"`

	// Token for OpenTSDB metrics access
	Token string `json:"token,omitempty"`
}