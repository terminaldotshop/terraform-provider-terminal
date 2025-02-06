// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package address

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/terminaldotshop/terraform-provider-terminal/internal/apijson"
)

type AddressModel struct {
	ID       types.String `tfsdk:"id" path:"id,optional"`
	City     types.String `tfsdk:"city" json:"city,required"`
	Country  types.String `tfsdk:"country" json:"country,required"`
	Name     types.String `tfsdk:"name" json:"name,required"`
	Street1  types.String `tfsdk:"street1" json:"street1,required"`
	Zip      types.String `tfsdk:"zip" json:"zip,required"`
	Phone    types.String `tfsdk:"phone" json:"phone,optional"`
	Province types.String `tfsdk:"province" json:"province,optional"`
	Street2  types.String `tfsdk:"street2" json:"street2,optional"`
	Data     types.String `tfsdk:"data" json:"data,computed"`
}

func (m AddressModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m AddressModel) MarshalJSONForUpdate(state AddressModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
