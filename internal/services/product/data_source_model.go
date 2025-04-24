// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package product

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/terminaldotshop/terraform-provider-terminal/internal/customfield"
)

type ProductDataSourceModel struct {
	ID   types.String                                         `tfsdk:"id" path:"id,required"`
	Data customfield.NestedObject[ProductDataDataSourceModel] `tfsdk:"data" json:"data,computed"`
}

type ProductDataDataSourceModel struct {
	ID           types.String                                                     `tfsdk:"id" json:"id,computed"`
	Description  types.String                                                     `tfsdk:"description" json:"description,computed"`
	Name         types.String                                                     `tfsdk:"name" json:"name,computed"`
	Variants     customfield.NestedObjectList[ProductDataVariantsDataSourceModel] `tfsdk:"variants" json:"variants,computed"`
	Order        types.Int64                                                      `tfsdk:"order" json:"order,computed"`
	Subscription types.String                                                     `tfsdk:"subscription" json:"subscription,computed"`
	Tags         customfield.NestedObject[ProductDataTagsDataSourceModel]         `tfsdk:"tags" json:"tags,computed"`
}

type ProductDataVariantsDataSourceModel struct {
	ID          types.String                                                     `tfsdk:"id" json:"id,computed"`
	Name        types.String                                                     `tfsdk:"name" json:"name,computed"`
	Price       types.Int64                                                      `tfsdk:"price" json:"price,computed"`
	Description types.String                                                     `tfsdk:"description" json:"description,computed"`
	Tags        customfield.NestedObject[ProductDataVariantsTagsDataSourceModel] `tfsdk:"tags" json:"tags,computed"`
}

type ProductDataVariantsTagsDataSourceModel struct {
	App          types.String `tfsdk:"app" json:"app,computed"`
	MarketEu     types.Bool   `tfsdk:"market_eu" json:"market_eu,computed"`
	MarketGlobal types.Bool   `tfsdk:"market_global" json:"market_global,computed"`
	MarketNa     types.Bool   `tfsdk:"market_na" json:"market_na,computed"`
}

type ProductDataTagsDataSourceModel struct {
	App          types.String `tfsdk:"app" json:"app,computed"`
	Color        types.String `tfsdk:"color" json:"color,computed"`
	Featured     types.Bool   `tfsdk:"featured" json:"featured,computed"`
	MarketEu     types.Bool   `tfsdk:"market_eu" json:"market_eu,computed"`
	MarketGlobal types.Bool   `tfsdk:"market_global" json:"market_global,computed"`
	MarketNa     types.Bool   `tfsdk:"market_na" json:"market_na,computed"`
}
