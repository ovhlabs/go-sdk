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

// DockerSLAveFrameworkAppVolume Application volumes
type DockerSLAveFrameworkAppVolume struct {

	// ContainerPath Container path
	ContainerPath int64 `json:"containerPath,omitempty"`

	// HostPath Host path
	HostPath int64 `json:"hostPath,omitempty"`

	// Mode Volume mode
	Mode string `json:"mode,omitempty"`
}
