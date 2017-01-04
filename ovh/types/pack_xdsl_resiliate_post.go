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

type PackXdslResiliatePost struct {

	ResiliationDate time.Time `json:"resiliationDate,omitempty"`

	ResiliationSurvey PackXdslResiliationSurvey `json:"resiliationSurvey,omitempty"`

	ServicesToKeep []float64 `json:"servicesToKeep,omitempty"`
}