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

// Customer answers for line diagnostic
type XdslLineDiagnosticSeltResult struct {

	// SELT test running date
	Date time.Time `json:"date,omitempty"`

	// Distance of the problem identified on the line (by SELT test), from NRA to customer
	Distance int64 `json:"distance,omitempty"`

	// Prelocalization of the problem
	Preloc string `json:"preloc,omitempty"`

	// Problem type identified by SELT test
	State string `json:"state,omitempty"`

	// SELT test status
	Status string `json:"status,omitempty"`
}