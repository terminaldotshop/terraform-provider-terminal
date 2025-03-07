// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package app

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/terminaldotshop/terraform-provider-terminal/internal/apijson"
	"github.com/terminaldotshop/terraform-provider-terminal/internal/customfield"
)

type AppModel struct {
	ID          types.String                           `tfsdk:"id" path:"id,optional"`
	Name        types.String                           `tfsdk:"name" json:"name,required"`
	RedirectUri types.String                           `tfsdk:"redirect_uri" json:"redirectURI,required"`
	Data        customfield.NestedObject[AppDataModel] `tfsdk:"data" json:"data,computed"`
}

func (m AppModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m AppModel) MarshalJSONForUpdate(state AppModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type AppDataModel struct {
	ID          types.String `tfsdk:"id" json:"id,computed"`
	Name        types.String `tfsdk:"name" json:"name,computed"`
	RedirectUri types.String `tfsdk:"redirect_uri" json:"redirectURI,computed"`
	Secret      types.String `tfsdk:"secret" json:"secret,computed"`
}
