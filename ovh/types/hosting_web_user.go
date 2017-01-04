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

// Hosting users
type HostingWebUser struct {

	// Home directory
	Home string `json:"home,omitempty"`

	// User IIS rights
	IisRemoteRights string `json:"iisRemoteRights,omitempty"`

	// Is this user primary
	IsPrimaryAccount bool `json:"isPrimaryAccount,omitempty"`

	// Login used to connect on FTP and SSH
	Login string `json:"login,omitempty"`

	// User ssh status
	SshState string `json:"sshState,omitempty"`

	// User status
	State string `json:"state,omitempty"`

	// User WebDav rights
	WebDavRights string `json:"webDavRights,omitempty"`
}