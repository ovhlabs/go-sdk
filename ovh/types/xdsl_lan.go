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

// LAN Configuration of the Modem
type XdslLan struct {

	// The IP address of the LAN interface of the modem
	IPAddress string `json:"IPAddress,omitempty"`

	// How the LAN interface of the modem is gettig its address
	AddressingType string `json:"addressingType,omitempty"`

	// Name of the LAN
	LanName string `json:"lanName,omitempty"`

	// The subnet mask of the LAN interface of the modem
	SubnetMask string `json:"subnetMask,omitempty"`

	// ID of the ongoing todo (NULL if none)
	TaskId int64 `json:"taskId,omitempty"`
}