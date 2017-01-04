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

type OrderDedicatedCloudUpgradeRessourcePost struct {

	UpgradeType string `json:"upgradeType,omitempty"`

	UpgradedRessourceId int64 `json:"upgradedRessourceId,omitempty"`

	UpgradedRessourceType string `json:"upgradedRessourceType,omitempty"`
}