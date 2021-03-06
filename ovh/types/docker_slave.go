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

import (
	"time"
)

// DockerSLAve A host on which containers can be deployed
type DockerSLAve struct {

	// CreatedAt Date of the resource creation
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// FlavorID The framework UUID
	FlavorID string `json:"flavorId,omitempty"`

	// ID The slave UUID
	ID string `json:"id,omitempty"`

	Metrics *DockerSLAveMetrics `json:"metrics,omitempty"`

	// Name The slave hostname
	Name string `json:"name,omitempty"`

	// Region The region where the slave is located
	Region string `json:"region,omitempty"`

	// Stack The stack to which the slave belongs
	Stack string `json:"stack,omitempty"`

	// State The state of the slave
	State string `json:"state,omitempty"`

	// UpdatedAt Date of the resource last update
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}
