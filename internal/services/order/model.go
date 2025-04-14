// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package order

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/terminaldotshop/terraform-provider-terminal/internal/apijson"
	"github.com/terminaldotshop/terraform-provider-terminal/internal/customfield"
)

type OrderModel struct {
	ID        types.String                             `tfsdk:"id" path:"id,optional"`
	AddressID types.String                             `tfsdk:"address_id" json:"addressID,required"`
	CardID    types.String                             `tfsdk:"card_id" json:"cardID,required"`
	Variants  *map[string]types.Int64                  `tfsdk:"variants" json:"variants,required"`
	Data      customfield.NestedObject[OrderDataModel] `tfsdk:"data" json:"data,computed"`
}

func (m OrderModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m OrderModel) MarshalJSONForUpdate(state OrderModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type OrderDataModel struct {
	ID       types.String                                      `tfsdk:"id" json:"id,computed"`
	Amount   customfield.NestedObject[OrderDataAmountModel]    `tfsdk:"amount" json:"amount,computed"`
	Created  types.String                                      `tfsdk:"created" json:"created,computed"`
	Items    customfield.NestedObjectList[OrderDataItemsModel] `tfsdk:"items" json:"items,computed"`
	Shipping customfield.NestedObject[OrderDataShippingModel]  `tfsdk:"shipping" json:"shipping,computed"`
	Tracking customfield.NestedObject[OrderDataTrackingModel]  `tfsdk:"tracking" json:"tracking,computed"`
	Index    types.Int64                                       `tfsdk:"index" json:"index,computed"`
}

type OrderDataAmountModel struct {
	Shipping types.Int64 `tfsdk:"shipping" json:"shipping,computed"`
	Subtotal types.Int64 `tfsdk:"subtotal" json:"subtotal,computed"`
}

type OrderDataItemsModel struct {
	ID               types.String `tfsdk:"id" json:"id,computed"`
	Amount           types.Int64  `tfsdk:"amount" json:"amount,computed"`
	Quantity         types.Int64  `tfsdk:"quantity" json:"quantity,computed"`
	Description      types.String `tfsdk:"description" json:"description,computed"`
	ProductVariantID types.String `tfsdk:"product_variant_id" json:"productVariantID,computed"`
}

type OrderDataShippingModel struct {
	City     types.String `tfsdk:"city" json:"city,computed"`
	Country  types.String `tfsdk:"country" json:"country,computed"`
	Name     types.String `tfsdk:"name" json:"name,computed"`
	Street1  types.String `tfsdk:"street1" json:"street1,computed"`
	Zip      types.String `tfsdk:"zip" json:"zip,computed"`
	Phone    types.String `tfsdk:"phone" json:"phone,computed"`
	Province types.String `tfsdk:"province" json:"province,computed"`
	Street2  types.String `tfsdk:"street2" json:"street2,computed"`
}

type OrderDataTrackingModel struct {
	Number          types.String `tfsdk:"number" json:"number,computed"`
	Service         types.String `tfsdk:"service" json:"service,computed"`
	Status          types.String `tfsdk:"status" json:"status,computed"`
	StatusDetails   types.String `tfsdk:"status_details" json:"statusDetails,computed"`
	StatusUpdatedAt types.String `tfsdk:"status_updated_at" json:"statusUpdatedAt,computed"`
	URL             types.String `tfsdk:"url" json:"url,computed"`
}
