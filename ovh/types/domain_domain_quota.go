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

// Values of quota account (pop, mailing list, redirection, responder and big pop)
type DomainDomainQuota struct {

	// Maximum number of mailboxes
	Account int64 `json:"account,omitempty"`

	// Maximum number of aliases
	Alias int64 `json:"alias,omitempty"`

	// Maximum number of mailing lists
	MailingList int64 `json:"mailingList,omitempty"`

	// Maximum number of redirections
	Redirection int64 `json:"redirection,omitempty"`

	// Maximum number of responders
	Responder int64 `json:"responder,omitempty"`
}