package service

import (
	"github.com/ovhlabs/go-sdk/ovh"
	"github.com/ovhlabs/go-sdk/ovh/ovhcli/common"

	"github.com/spf13/cobra"
)

var cmdInfo = &cobra.Command{
	Use:   "info",
	Short: "Get Application Info: ovhcli dbaas queue service info (--name=AppName | <--id=appID>)",
	Run: func(cmd *cobra.Command, args []string) {

		client, err := ovh.NewDefaultClient()
		common.Check(err)

		if name != "" {
			app, errInfo := client.DBaasQueueAppInfoByName(name)
			common.Check(errInfo)
			common.FormatOutputDef(app)
		} else {
			app, errInfo := client.DBaasQueueAppInfo(id)
			common.Check(errInfo)
			common.FormatOutputDef(app)
		}

	},
}
