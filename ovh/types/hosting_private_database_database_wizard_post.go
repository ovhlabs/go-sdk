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

type HostingPrivateDatabaseDatabaseWizardPost struct {

	DatabaseName string `json:"databaseName,omitempty"`

	Grant string `json:"grant,omitempty"`

	Password string `json:"password,omitempty"`

	UserName string `json:"userName,omitempty"`
}