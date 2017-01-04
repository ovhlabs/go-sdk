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

// All data needed to use a payment mean
type BillingOrderPaymentMean struct {

	Fee float64 `json:"fee,omitempty"`

	HtmlForm string `json:"htmlForm,omitempty"`

	HttpMethod string `json:"httpMethod,omitempty"`

	Logo string `json:"logo,omitempty"`

	Parameters []BillingOrderPaymentMeanHttpParameter `json:"parameters,omitempty"`

	SubType string `json:"subType,omitempty"`

	Url string `json:"url,omitempty"`
}