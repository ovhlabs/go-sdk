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

// RouterNetwork Network
type RouterNetwork struct {
	CreationDate *time.Time `json:"creationDate,omitempty"`

	Description string `json:"description,omitempty"`

	ID int64 `json:"id,omitempty"`

	// IPNet Gateway IP / CIDR Netmask
	IPNet string `json:"ipNet,omitempty"`

	Status string `json:"status,omitempty"`

	VlanTag int64 `json:"vlanTag,omitempty"`
}
