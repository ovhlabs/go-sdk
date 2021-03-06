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

// HorizonViewStorage Cloud Desktop Infrastructure storages
type HorizonViewStorage struct {

	// FullProfile Human-Readable profile name
	FullProfile string `json:"fullProfile,omitempty"`

	// ID Storage ids
	ID int64 `json:"id,omitempty"`

	// Name Filer name
	Name string `json:"name,omitempty"`

	// Profile Commercial profile name
	Profile string `json:"profile,omitempty"`

	Size *DedicatedCloudFilerSize `json:"size,omitempty"`

	// SpaceFree Available space of this datastore, in GB
	SpaceFree float64 `json:"spaceFree,omitempty"`

	// SpaceProvisionned Provisionned space of this datastore, in GB
	SpaceProvisionned float64 `json:"spaceProvisionned,omitempty"`

	// SpaceUsed Used Space of this filer, in GB
	SpaceUsed float64 `json:"spaceUsed,omitempty"`

	// State State of the filer
	State string `json:"state,omitempty"`

	// VMTotal Number of virtual machine on the filer
	VMTotal int64 `json:"vmTotal,omitempty"`
}
