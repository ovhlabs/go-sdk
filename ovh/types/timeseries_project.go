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

// TimeseriesProject Timeseries project
type TimeseriesProject struct {

	// Description description of your project
	Description string `json:"description,omitempty"`

	// DisplayName name of your project
	DisplayName string `json:"displayName,omitempty"`

	// OfferID subscribed offer
	OfferID string `json:"offerId,omitempty"`

	// RegionID region where your data are located
	RegionID string `json:"regionId,omitempty"`

	// ServiceName timeseries Project id
	ServiceName string `json:"serviceName,omitempty"`

	// Status project status
	Status string `json:"status,omitempty"`
}
