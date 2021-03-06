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

// SmsPhonebook Phone book
type SmsPhonebook struct {

	// BookKey Identifier of the phonebook
	BookKey string `json:"bookKey,omitempty"`

	// IsReadonly Set if phonebook is readonly
	IsReadonly bool `json:"isReadonly,omitempty"`

	// Name Phonebook name
	Name string `json:"name,omitempty"`

	// PhoneKey Phone key identifier between the phone and phonebooks
	PhoneKey string `json:"phoneKey,omitempty"`
}
