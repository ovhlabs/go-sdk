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

// License usage statistics.
type LicenseOfficeStatistics struct {

	// Date of the statistics.
	Date time.Time `json:"date,omitempty"`

	// List of lines associated to this statistics entity.
	Lines []LicenseOfficeStatisticsLine `json:"lines,omitempty"`
}