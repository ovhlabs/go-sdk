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

// Attributes of the slave flavor
type DockerSlaveFlavor struct {

	// The network bandwidth, in Mbps
	Bandwidth int64 `json:"bandwidth,omitempty"`

	// The amount of (v)CPUs
	Cpus int64 `json:"cpus,omitempty"`

	// The disk size, in GB
	Disk int64 `json:"disk,omitempty"`

	// Wether the disk is HA (stored in Ceph) or local (SSD)
	DiskHa bool `json:"diskHa,omitempty"`

	// The flavor UUID
	Id string `json:"id,omitempty"`

	// Whether the flavor is an Openstack or dedicated flavor
	IsVm bool `json:"isVm,omitempty"`

	// The amount of RAM, in MB
	Ram int64 `json:"ram,omitempty"`
}