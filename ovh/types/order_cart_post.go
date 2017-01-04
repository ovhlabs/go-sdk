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

type OrderCartPost struct {

	Description string `json:"description,omitempty"`

	Expire time.Time `json:"expire,omitempty"`

	OvhSubsidiary string `json:"ovhSubsidiary,omitempty"`
}