// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package subscription

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/terminaldotshop/terraform-provider-terminal/internal/apijson"
	"github.com/terminaldotshop/terraform-provider-terminal/internal/customfield"
)

type SubscriptionModel struct {
	ID               types.String                                        `tfsdk:"id" json:"id,required"`
	Created          types.String                                        `tfsdk:"created" json:"created,required"`
	Price            types.Int64                                         `tfsdk:"price" json:"price,required"`
	ProductVariantID types.String                                        `tfsdk:"product_variant_id" json:"productVariantID,required"`
	Quantity         types.Int64                                         `tfsdk:"quantity" json:"quantity,required"`
	Next             types.String                                        `tfsdk:"next" json:"next,optional"`
	AddressID        types.String                                        `tfsdk:"address_id" json:"addressID,required"`
	CardID           types.String                                        `tfsdk:"card_id" json:"cardID,required"`
	Schedule         customfield.NestedObject[SubscriptionScheduleModel] `tfsdk:"schedule" json:"schedule,computed_optional"`
	Data             customfield.NestedObject[SubscriptionDataModel]     `tfsdk:"data" json:"data,computed"`
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

type SubscriptionDataModel struct {
	ID               types.String                                            `tfsdk:"id" json:"id,computed"`
	AddressID        types.String                                            `tfsdk:"address_id" json:"addressID,computed"`
	CardID           types.String                                            `tfsdk:"card_id" json:"cardID,computed"`
	Created          types.String                                            `tfsdk:"created" json:"created,computed"`
	Price            types.Int64                                             `tfsdk:"price" json:"price,computed"`
	ProductVariantID types.String                                            `tfsdk:"product_variant_id" json:"productVariantID,computed"`
	Quantity         types.Int64                                             `tfsdk:"quantity" json:"quantity,computed"`
	Next             types.String                                            `tfsdk:"next" json:"next,computed"`
	Schedule         customfield.NestedObject[SubscriptionDataScheduleModel] `tfsdk:"schedule" json:"schedule,computed"`
}

type SubscriptionDataScheduleModel struct {
	Type     types.String `tfsdk:"type" json:"type,computed"`
	Interval types.Int64  `tfsdk:"interval" json:"interval,computed"`
}
