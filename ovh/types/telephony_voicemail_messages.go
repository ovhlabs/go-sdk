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

// Voicemail message
type TelephonyVoicemailMessages struct {

	// Callee number in international format
	Callee string `json:"callee,omitempty"`

	// Caller number in international format
	Caller string `json:"caller,omitempty"`

	// Message datetime creation
	CreationDatetime time.Time `json:"creationDatetime,omitempty"`

	// Voicemessage directory
	Dir string `json:"dir,omitempty"`

	// Message duration (in seconds)
	Duration int64 `json:"duration,omitempty"`

	// Uniq voicemail message identifier
	Id int64 `json:"id,omitempty"`
}