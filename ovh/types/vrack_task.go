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

// vrack tasks
type VrackTask struct {

	Function string `json:"function,omitempty"`

	Id int64 `json:"id,omitempty"`

	LastUpdate time.Time `json:"lastUpdate,omitempty"`

	OrderId int64 `json:"orderId,omitempty"`

	ServiceName string `json:"serviceName,omitempty"`

	Status string `json:"status,omitempty"`

	TargetDomain string `json:"targetDomain,omitempty"`

	TodoDate time.Time `json:"todoDate,omitempty"`
}