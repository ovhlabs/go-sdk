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

type SmsUsersReceiversPost struct {

	AutoUpdate bool `json:"autoUpdate,omitempty"`

	CsvUrl string `json:"csvUrl,omitempty"`

	Description string `json:"description,omitempty"`

	SlotId int64 `json:"slotId,omitempty"`
}