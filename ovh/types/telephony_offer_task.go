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

// Operation on a telephony offer
type TelephonyOfferTask struct {

	// Actual action that will be executed
	Action string `json:"action,omitempty"`

	// Planned execution date
	ExecutionDate time.Time `json:"executionDate,omitempty"`

	// Current status of the task
	Status string `json:"status,omitempty"`

	TaskId int64 `json:"taskId,omitempty"`

	// Type of operation that will be executed
	Type_ string `json:"type,omitempty"`
}