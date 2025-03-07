// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package address

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/terminaldotshop/terraform-provider-terminal/internal/customfield"
)

type AddressDataSourceModel struct {
	ID   types.String                                         `tfsdk:"id" path:"id,required"`
	Data customfield.NestedObject[AddressDataDataSourceModel] `tfsdk:"data" json:"data,computed"`
}

type AddressDataDataSourceModel struct {
	ID       types.String `tfsdk:"id" json:"id,computed"`
	City     types.String `tfsdk:"city" json:"city,computed"`
	Country  types.String `tfsdk:"country" json:"country,computed"`
	Name     types.String `tfsdk:"name" json:"name,computed"`
	Street1  types.String `tfsdk:"street1" json:"street1,computed"`
	Zip      types.String `tfsdk:"zip" json:"zip,computed"`
	Phone    types.String `tfsdk:"phone" json:"phone,computed"`
	Province types.String `tfsdk:"province" json:"province,computed"`
	Street2  types.String `tfsdk:"street2" json:"street2,computed"`
}
