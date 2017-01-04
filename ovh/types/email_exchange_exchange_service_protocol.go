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

// Protocol access policy for this Exchange service
type EmailExchangeExchangeServiceProtocol struct {

	// IMAP protocol enabled on this Exchange service
	IMAP bool `json:"IMAP,omitempty"`

	// POP protocol enabled on this Exchange service
	POP bool `json:"POP,omitempty"`

	// ActiveSync protocol enabled on this Exchange service
	ActiveSync bool `json:"activeSync,omitempty"`

	// ActiveSync policy to apply at device's first connection
	ActiveSyncPolicy string `json:"activeSyncPolicy,omitempty"`

	// Creation date
	CreationDate time.Time `json:"creationDate,omitempty"`

	// Last update date
	LastUpdate time.Time `json:"lastUpdate,omitempty"`

	// Pending task id
	TaskPendingId int64 `json:"taskPendingId,omitempty"`

	// Web mail protocol enabled on this Exchange service
	WebMail bool `json:"webMail,omitempty"`
}