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

// XdslAsyncTaskPackXdslMigrationOfferResponse Async task
type XdslAsyncTaskPackXdslMigrationOfferResponse struct {

	// TError Error
	TError string `json:"error,omitempty"`

	Result *PackXdslMigrationOfferResponse `json:"result,omitempty"`

	// Status Status of the call
	Status string `json:"status,omitempty"`
}
