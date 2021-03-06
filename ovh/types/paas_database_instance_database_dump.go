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

// PaasDatabaseInstanceDatabaseDump Database Dumps
type PaasDatabaseInstanceDatabaseDump struct {

	// CreationDate Creation date of the dump
	CreationDate *time.Time `json:"creationDate,omitempty"`

	// DatabaseName Database name associated to this dump
	DatabaseName string `json:"databaseName,omitempty"`

	// DumpID Dump id
	DumpID string `json:"dumpId,omitempty"`

	// Email email
	Email string `json:"email,omitempty"`

	// ExpirationDate Deletion date of the dump
	ExpirationDate *time.Time `json:"expirationDate,omitempty"`

	// InstanceID Instance id associated to this dump
	InstanceID string `json:"instanceId,omitempty"`

	// Name Dump name
	Name string `json:"name,omitempty"`

	// Status Dump status
	Status string `json:"status,omitempty"`

	// TaskIDs The task id working on this object
	TaskIDs []string `json:"taskIds,omitempty"`

	// URL Dump url
	URL string `json:"url,omitempty"`
}
