// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package address

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/terminaldotshop/terraform-provider-terminal/internal/apijson"
	"github.com/terminaldotshop/terraform-provider-terminal/internal/customfield"
)

type AddressModel struct {
	ID       types.String                               `tfsdk:"id" path:"id,optional"`
	City     types.String                               `tfsdk:"city" json:"city,required,no_refresh"`
	Country  types.String                               `tfsdk:"country" json:"country,required,no_refresh"`
	Name     types.String                               `tfsdk:"name" json:"name,required,no_refresh"`
	Street1  types.String                               `tfsdk:"street1" json:"street1,required,no_refresh"`
	Zip      types.String                               `tfsdk:"zip" json:"zip,required,no_refresh"`
	Phone    types.String                               `tfsdk:"phone" json:"phone,optional,no_refresh"`
	Province types.String                               `tfsdk:"province" json:"province,optional,no_refresh"`
	Street2  types.String                               `tfsdk:"street2" json:"street2,optional,no_refresh"`
	Data     customfield.NestedObject[AddressDataModel] `tfsdk:"data" json:"data,computed"`
}

func (m AddressModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m AddressModel) MarshalJSONForUpdate(state AddressModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type AddressDataModel struct {
	ID       types.String `tfsdk:"id" json:"id,computed"`
	City     types.String `tfsdk:"city" json:"city,computed"`
	Country  types.String `tfsdk:"country" json:"country,computed"`
	Created  types.String `tfsdk:"created" json:"created,computed"`
	Name     types.String `tfsdk:"name" json:"name,computed"`
	Street1  types.String `tfsdk:"street1" json:"street1,computed"`
	Zip      types.String `tfsdk:"zip" json:"zip,computed"`
	Phone    types.String `tfsdk:"phone" json:"phone,computed"`
	Province types.String `tfsdk:"province" json:"province,computed"`
	Street2  types.String `tfsdk:"street2" json:"street2,computed"`
}
