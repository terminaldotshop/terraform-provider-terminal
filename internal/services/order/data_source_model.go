// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package order

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/terminaldotshop/terraform-provider-terminal/internal/customfield"
)

type OrderDataSourceModel struct {
	ID   types.String                                       `tfsdk:"id" path:"id,required"`
	Data customfield.NestedObject[OrderDataDataSourceModel] `tfsdk:"data" json:"data,computed"`
}

type OrderDataDataSourceModel struct {
	ID       types.String                                                `tfsdk:"id" json:"id,computed"`
	Amount   customfield.NestedObject[OrderDataAmountDataSourceModel]    `tfsdk:"amount" json:"amount,computed"`
	Items    customfield.NestedObjectList[OrderDataItemsDataSourceModel] `tfsdk:"items" json:"items,computed"`
	Shipping customfield.NestedObject[OrderDataShippingDataSourceModel]  `tfsdk:"shipping" json:"shipping,computed"`
	Tracking customfield.NestedObject[OrderDataTrackingDataSourceModel]  `tfsdk:"tracking" json:"tracking,computed"`
	Index    types.Int64                                                 `tfsdk:"index" json:"index,computed"`
}

type OrderDataAmountDataSourceModel struct {
	Shipping types.Int64 `tfsdk:"shipping" json:"shipping,computed"`
	Subtotal types.Int64 `tfsdk:"subtotal" json:"subtotal,computed"`
}

type OrderDataItemsDataSourceModel struct {
	ID               types.String `tfsdk:"id" json:"id,computed"`
	Amount           types.Int64  `tfsdk:"amount" json:"amount,computed"`
	Quantity         types.Int64  `tfsdk:"quantity" json:"quantity,computed"`
	Description      types.String `tfsdk:"description" json:"description,computed"`
	ProductVariantID types.String `tfsdk:"product_variant_id" json:"productVariantID,computed"`
}

type OrderDataShippingDataSourceModel struct {
	City     types.String `tfsdk:"city" json:"city,computed"`
	Country  types.String `tfsdk:"country" json:"country,computed"`
	Name     types.String `tfsdk:"name" json:"name,computed"`
	Street1  types.String `tfsdk:"street1" json:"street1,computed"`
	Zip      types.String `tfsdk:"zip" json:"zip,computed"`
	Phone    types.String `tfsdk:"phone" json:"phone,computed"`
	Province types.String `tfsdk:"province" json:"province,computed"`
	Street2  types.String `tfsdk:"street2" json:"street2,computed"`
}

type OrderDataTrackingDataSourceModel struct {
	Number          types.String `tfsdk:"number" json:"number,computed"`
	Service         types.String `tfsdk:"service" json:"service,computed"`
	Status          types.String `tfsdk:"status" json:"status,computed"`
	StatusDetails   types.String `tfsdk:"status_details" json:"statusDetails,computed"`
	StatusUpdatedAt types.String `tfsdk:"status_updated_at" json:"statusUpdatedAt,computed"`
	URL             types.String `tfsdk:"url" json:"url,computed"`
}
