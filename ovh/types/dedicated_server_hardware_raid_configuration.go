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

// Simulated result of how the hardware RAID template will be configured on this server
type DedicatedServerHardwareRaidConfiguration struct {

	Capacity DedicatedServerHardwareRaidConfigurationCapacity `json:"capacity,omitempty"`

	// Disk count
	DiskCount int64 `json:"diskCount,omitempty"`

	DiskSize DedicatedServerHardwareRaidConfigurationDiskSize `json:"diskSize,omitempty"`

	// Size of disk spans on RAID
	DiskSpanSize int64 `json:"diskSpanSize,omitempty"`

	// RAID mode
	Mode string `json:"mode,omitempty"`

	// RAID configuration name
	Name string `json:"name,omitempty"`

	// RAID controller type
	Type_ string `json:"type,omitempty"`
}