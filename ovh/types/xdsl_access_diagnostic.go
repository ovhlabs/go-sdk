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

// Diagnostic of the access
type XdslAccessDiagnostic struct {

	Capabilities XdslAccessDiagnosticCapabilities `json:"capabilities,omitempty"`

	// Datime of the diagnostic
	DiagnosticTime time.Time `json:"diagnosticTime,omitempty"`

	// Is there an ongoing genericIncident on the access ?
	Incident bool `json:"incident,omitempty"`

	// Is the access active on its primary or secondary LNS
	IsActiveOnLns bool `json:"isActiveOnLns,omitempty"`

	// Is the modem connected ?
	IsModemConnected bool `json:"isModemConnected,omitempty"`

	// Test details by line
	LineDetails []XdslLineDiagnostic `json:"lineDetails,omitempty"`

	// Does the access ping ?
	Ping bool `json:"ping,omitempty"`

	// Remaining number of diagnostic for this access
	Remaining int64 `json:"remaining,omitempty"`
}