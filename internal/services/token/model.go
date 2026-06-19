// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package token

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/terminaldotshop/terraform-provider-terminal/internal/apijson"
	"github.com/terminaldotshop/terraform-provider-terminal/internal/customfield"
)

type TokenModel struct {
	ID   types.String                             `tfsdk:"id" path:"id,optional"`
	Data customfield.NestedObject[TokenDataModel] `tfsdk:"data" json:"data,computed"`
}

func (m TokenModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m TokenModel) MarshalJSONForUpdate(state TokenModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type TokenDataModel struct {
	ID      types.String `tfsdk:"id" json:"id,computed"`
	Token   types.String `tfsdk:"token" json:"token,computed"`
	Created types.String `tfsdk:"created" json:"created,computed"`
}
