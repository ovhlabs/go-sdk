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

// Details about the resiliation
type XdslResiliationFollowUpDetail struct {

	// Date when the resiliation will take effect
	DateTodo time.Time `json:"dateTodo,omitempty"`

	// If the customer needs to return his modem
	NeedModemReturn bool `json:"needModemReturn,omitempty"`

	// Date when the resiliation was done
	RegistrationDate time.Time `json:"registrationDate,omitempty"`

	// Status of the resiliation
	Status string `json:"status,omitempty"`
}