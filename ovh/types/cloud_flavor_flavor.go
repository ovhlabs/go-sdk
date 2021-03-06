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

// CloudFlavor Flavor
type CloudFlavor struct {

	// Available Available in stock
	Available bool `json:"available,omitempty"`

	// Disk Number of disks
	Disk int64 `json:"disk,omitempty"`

	// ID Flavor id
	ID string `json:"id,omitempty"`

	// InboundBandwidth Max capacity of inbound traffic in Mbit/s
	InboundBandwidth int64 `json:"inboundBandwidth,omitempty"`

	// Name Flavor name
	Name string `json:"name,omitempty"`

	// OsType OS to install on
	OsType string `json:"osType,omitempty"`

	// OutboundBandwidth Max capacity of outbound traffic in Mbit/s
	OutboundBandwidth int64 `json:"outboundBandwidth,omitempty"`

	// RAM Ram quantity (Gio)
	RAM int64 `json:"ram,omitempty"`

	// Region Flavor region
	Region string `json:"region,omitempty"`

	// TType Flavor type
	TType string `json:"type,omitempty"`

	// Vcpus Number of VCPUs
	Vcpus int64 `json:"vcpus,omitempty"`
}
