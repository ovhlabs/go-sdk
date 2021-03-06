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

// DedicatedServerRtmMemory A structure describing informations about server memory
type DedicatedServerRtmMemory struct {
	Capacity *DedicatedServerRtmMemoryCapacity `json:"capacity,omitempty"`

	// Slot Memory slot
	Slot string `json:"slot,omitempty"`
}
