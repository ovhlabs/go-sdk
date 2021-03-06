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

// RegistryTag An image tag
type RegistryTag struct {

	// CreatedAt Date of the resource creation
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// ID The tag id
	ID string `json:"id,omitempty"`

	// Name The tag name
	Name string `json:"name,omitempty"`

	// UpdatedAt Date of the resource last update
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}
