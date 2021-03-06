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

// OrderTelephonyPortabilityPost ...
type OrderTelephonyPortabilityPost struct {
	Building string `json:"building,omitempty"`

	CallNumber string `json:"callNumber,omitempty"`

	City string `json:"city,omitempty"`

	Comment string `json:"comment,omitempty"`

	ContactName string `json:"contactName,omitempty"`

	ContactNumber string `json:"contactNumber,omitempty"`

	Country string `json:"country,omitempty"`

	DesireDate *time.Time `json:"desireDate,omitempty"`

	DisplayUniversalDirectory bool `json:"displayUniversalDirectory,omitempty"`

	Door string `json:"door,omitempty"`

	FirstName string `json:"firstName,omitempty"`

	Floor float64 `json:"floor,omitempty"`

	GroupNumber string `json:"groupNumber,omitempty"`

	LineToRedirectAliasTo string `json:"lineToRedirectAliasTo,omitempty"`

	ListNumbers []string `json:"listNumbers,omitempty"`

	Name string `json:"name,omitempty"`

	Offer string `json:"offer,omitempty"`

	Rio string `json:"rio,omitempty"`

	Siret string `json:"siret,omitempty"`

	SocialReason string `json:"socialReason,omitempty"`

	Stair float64 `json:"stair,omitempty"`

	StreetName string `json:"streetName,omitempty"`

	StreetNumber float64 `json:"streetNumber,omitempty"`

	StreetNumberExtra string `json:"streetNumberExtra,omitempty"`

	StreetType string `json:"streetType,omitempty"`

	TType string `json:"type,omitempty"`

	Zip string `json:"zip,omitempty"`
}
