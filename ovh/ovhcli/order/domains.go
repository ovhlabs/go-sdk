package order

import (
	"github.com/ovhlabs/go-sdk/ovh"
	"github.com/ovhlabs/go-sdk/ovh/ovhcli/common"
	"github.com/ovhlabs/go-sdk/ovh/types"
	"github.com/spf13/cobra"
)

var withOffer string
var withConfigs string

var withOwner string

func init() {
	CmdDomain.PersistentFlags().StringVarP(&withOffer, "withOffer", "", "gold", "offer on your domain (gold, diamond, platinium)")
	CmdDomain.PersistentFlags().StringVarP(&withConfigs, "withConfigs", "", "", "configs file")

	CmdDomainTrade.PersistentFlags().StringVarP(&withOwner, "withOwner", "", "", "configs file")

}

// CmdDomain order domain
var CmdDomain = &cobra.Command{
	Use:   "domain <domain>",
	Short: "Order domain",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			common.WrongUsage(cmd)
		}
		domain := args[0]

		client, err := ovh.NewDefaultClient()
		common.Check(err)

		cart, err := client.OrderCreateCart(types.OrderCartPost{OvhSubsidiary: "FR"})
		common.Check(err)

		err = client.OrderAssignCart(cart.CartID)
		common.Check(err)

		_, err = client.OrderAddProductDomain(cart.CartID, types.OrderCartDomainPost{
			Domain: domain,
		})
		common.Check(err)

		order, err := client.OrderPostCheckoutCart(cart.CartID, false)
		common.Check(err)
		common.FormatOutputDef(order)

	},
}

// CmdDomainTransfer in order domain trade
var CmdDomainTransfer = &cobra.Command{
	Use:   "domainTransfer <domain>",
	Short: "Order domain transfer",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			common.WrongUsage(cmd)
		}
		domain := args[0]

		client, err := ovh.NewDefaultClient()
		common.Check(err)

		cart, err := client.OrderCreateCart(types.OrderCartPost{OvhSubsidiary: "FR"})
		common.Check(err)

		err = client.OrderAssignCart(cart.CartID)
		common.Check(err)

		_, err = client.OrderAddProductDomainTransfer(cart.CartID, types.OrderCartDomainTransferPost{
			Domain: domain,
		})
		common.Check(err)

		order, err := client.OrderPostCheckoutCart(cart.CartID, false)
		common.Check(err)
		common.FormatOutputDef(order)
	},
}

// CmdDomainTrade in order domain trade
var CmdDomainTrade = &cobra.Command{
	Use:   "domainTrade <domain>",
	Short: "Order domain trade",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			common.WrongUsage(cmd)
		}
		domain := args[0]

		if withOwner == "" {
			common.Exit("Missing withOwner option")
		}

		client, err := ovh.NewDefaultClient()
		common.Check(err)

		cart, err := client.OrderCreateCart(types.OrderCartPost{OvhSubsidiary: "FR"})
		common.Check(err)

		err = client.OrderAssignCart(cart.CartID)
		common.Check(err)

		serviceOptions, err := client.OrderGetCartServiceOptions(domain)
		common.Check(err)

		var serviceOptionChoosed *types.OrderCartGenericOptionDefinition
		for i := range serviceOptions {
			serviceOption := serviceOptions[i]
			if serviceOption.Family == "trade" {
				serviceOptionChoosed = &serviceOption
				break
			}
		}
		if serviceOptionChoosed == nil {
			common.Exit("Cannot find service options of type trade for domain %s", domain)
		}

		item, err := client.OrderAddCartServiceOption(domain, types.OrderCartServiceOptionDomainPost{
			Duration:    "P1Y", //always P1Y for a trade
			Quantity:    1,     //always 1 for quantity for a trade
			CartID:      cart.CartID,
			PlanCode:    serviceOptionChoosed.PlanCode,
			PricingMode: "default",
		})
		common.Check(err)

		_, err = client.OrderCartAddConfiguration(cart.CartID, item.ItemID, "OWNER_CONTACT", withOwner)
		common.Check(err)

		order, err := client.OrderPostCheckoutCart(cart.CartID, true)
		common.Check(err)

		common.FormatOutputDef(order)

	},
}
