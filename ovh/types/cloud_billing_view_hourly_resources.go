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

// HourlyResources
type CloudBillingViewHourlyResources struct {

	// Details about hourly instances
	Instance []CloudBillingViewHourlyInstance `json:"instance,omitempty"`

	// Details about hourly snapshots
	Snapshot []CloudBillingViewHourlySnapshot `json:"snapshot,omitempty"`

	// Details about hourly storage
	Storage []CloudBillingViewHourlyStorage `json:"storage,omitempty"`

	// Details about hourly volumes
	Volume []CloudBillingViewHourlyVolume `json:"volume,omitempty"`
}