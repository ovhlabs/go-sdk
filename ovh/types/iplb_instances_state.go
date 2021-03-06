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

// IPLBInstancesState A structure describing the current state of an IPLB instances
type IPLBInstancesState struct {

	// InternalID Internal ID of this IPLB instance
	InternalID int64 `json:"internalId,omitempty"`

	// LastUpdateDate Last update date
	LastUpdateDate *time.Time `json:"lastUpdateDate,omitempty"`

	// State Current state of this IPLB instance
	State string `json:"state,omitempty"`

	// Zone zone of this IPLB instance
	Zone string `json:"zone,omitempty"`
}
