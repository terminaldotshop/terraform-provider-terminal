// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package subscription

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/terminaldotshop/terraform-provider-terminal/internal/apijson"
	"github.com/terminaldotshop/terraform-provider-terminal/internal/customfield"
)

type SubscriptionModel struct {
	ID               types.String                                        `tfsdk:"id" json:"id,required"`
	AddressID        types.String                                        `tfsdk:"address_id" json:"addressID,required"`
	CardID           types.String                                        `tfsdk:"card_id" json:"cardID,required"`
	ProductVariantID types.String                                        `tfsdk:"product_variant_id" json:"productVariantID,required"`
	Quantity         types.Int64                                         `tfsdk:"quantity" json:"quantity,required"`
	Next             types.String                                        `tfsdk:"next" json:"next,optional"`
	Schedule         customfield.NestedObject[SubscriptionScheduleModel] `tfsdk:"schedule" json:"schedule,computed_optional"`
	Data             types.String                                        `tfsdk:"data" json:"data,computed"`
}

func (m SubscriptionModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m SubscriptionModel) MarshalJSONForUpdate(state SubscriptionModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type SubscriptionScheduleModel struct {
	Type     types.String `tfsdk:"type" json:"type,required"`
	Interval types.Int64  `tfsdk:"interval" json:"interval,optional"`
}
