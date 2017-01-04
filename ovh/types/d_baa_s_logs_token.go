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

// Token
type DBaaSLogsToken struct {

	// Token creation
	CreatedAt time.Time `json:"createdAt,omitempty"`

	// Token name
	Name string `json:"name,omitempty"`

	// Token UUID
	TokenId string `json:"tokenId,omitempty"`

	// Token last update
	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	// Token value
	Value string `json:"value,omitempty"`
}