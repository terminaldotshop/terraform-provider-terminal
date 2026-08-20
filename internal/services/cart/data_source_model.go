// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cart

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/terminaldotshop/terraform-provider-terminal/internal/customfield"
)

type CartDataSourceModel struct {
	Data customfield.NestedObject[CartDataDataSourceModel] `tfsdk:"data" json:"data,computed"`
}

type CartDataDataSourceModel struct {
	Amount    customfield.NestedObject[CartDataAmountDataSourceModel]    `tfsdk:"amount" json:"amount,computed"`
	Items     customfield.NestedObjectList[CartDataItemsDataSourceModel] `tfsdk:"items" json:"items,computed"`
	Subtotal  types.Int64                                                `tfsdk:"subtotal" json:"subtotal,computed"`
	AddressID types.String                                               `tfsdk:"address_id" json:"addressID,computed"`
	CardID    types.String                                               `tfsdk:"card_id" json:"cardID,computed"`
	Shipping  customfield.NestedObject[CartDataShippingDataSourceModel]  `tfsdk:"shipping" json:"shipping,computed"`
}

type CartDataAmountDataSourceModel struct {
	Subtotal types.Int64 `tfsdk:"subtotal" json:"subtotal,computed"`
	Shipping types.Int64 `tfsdk:"shipping" json:"shipping,computed"`
	Total    types.Int64 `tfsdk:"total" json:"total,computed"`
}

type CartDataItemsDataSourceModel struct {
	ID               types.String `tfsdk:"id" json:"id,computed"`
	ProductVariantID types.String `tfsdk:"product_variant_id" json:"productVariantID,computed"`
	Quantity         types.Int64  `tfsdk:"quantity" json:"quantity,computed"`
	Subtotal         types.Int64  `tfsdk:"subtotal" json:"subtotal,computed"`
}

type CartDataShippingDataSourceModel struct {
	Service   types.String `tfsdk:"service" json:"service,computed"`
	Timeframe types.String `tfsdk:"timeframe" json:"timeframe,computed"`
}
