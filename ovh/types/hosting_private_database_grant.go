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

// Grants
type HostingPrivateDatabaseGrant struct {

	// Creation date of the grant
	CreationDate time.Time `json:"creationDate,omitempty"`

	// Database name where grant is set
	DatabaseName string `json:"databaseName,omitempty"`

	// Grant set
	Grant string `json:"grant,omitempty"`
}