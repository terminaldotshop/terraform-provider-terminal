// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package subscription

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/terminaldotshop/terraform-provider-terminal/internal/customfield"
)

type SubscriptionDataSourceModel struct {
	ID   types.String                                              `tfsdk:"id" path:"id,required"`
	Data customfield.NestedObject[SubscriptionDataDataSourceModel] `tfsdk:"data" json:"data,computed"`
}

type SubscriptionDataDataSourceModel struct {
	ID               types.String                                                      `tfsdk:"id" json:"id,computed"`
	AddressID        types.String                                                      `tfsdk:"address_id" json:"addressID,computed"`
	CardID           types.String                                                      `tfsdk:"card_id" json:"cardID,computed"`
	Created          types.String                                                      `tfsdk:"created" json:"created,computed"`
	ProductVariantID types.String                                                      `tfsdk:"product_variant_id" json:"productVariantID,computed"`
	Quantity         types.Int64                                                       `tfsdk:"quantity" json:"quantity,computed"`
	Next             types.String                                                      `tfsdk:"next" json:"next,computed"`
	Schedule         customfield.NestedObject[SubscriptionDataScheduleDataSourceModel] `tfsdk:"schedule" json:"schedule,computed"`
}

type SubscriptionDataScheduleDataSourceModel struct {
	Type     types.String `tfsdk:"type" json:"type,computed"`
	Interval types.Int64  `tfsdk:"interval" json:"interval,computed"`
}
