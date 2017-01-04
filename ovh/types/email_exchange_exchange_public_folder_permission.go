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

// Exchange organization public folder permission
type EmailExchangeExchangePublicFolderPermission struct {

	// Access right set for the account
	AccessRights string `json:"accessRights,omitempty"`

	// Account id
	AllowedAccountId int64 `json:"allowedAccountId,omitempty"`

	// Creation date
	CreationDate time.Time `json:"creationDate,omitempty"`

	State string `json:"state,omitempty"`

	// task pending id
	TaskPendingId int64 `json:"taskPendingId,omitempty"`
}