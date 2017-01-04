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

// Mailing List
type EmailDomainMailingList struct {

	// Id of mailing list
	Id int64 `json:"id,omitempty"`

	// Language of mailing list
	Language string `json:"language,omitempty"`

	// Name of mailing list
	Name string `json:"name,omitempty"`

	// Subscribers number of mailing list
	NbSubscribers int64 `json:"nbSubscribers,omitempty"`

	// Last update subscribers
	NbSubscribersUpdateDate time.Time `json:"nbSubscribersUpdateDate,omitempty"`

	Options DomainDomainMlOptionsStruct `json:"options,omitempty"`

	// Owner email of mailing list
	OwnerEmail string `json:"ownerEmail,omitempty"`

	// Email to reply of mailing list
	ReplyTo string `json:"replyTo,omitempty"`
}