// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package token

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/terminaldotshop/terraform-provider-terminal/internal/customfield"
)

type TokenDataSourceModel struct {
	ID   types.String                                       `tfsdk:"id" path:"id,required"`
	Data customfield.NestedObject[TokenDataDataSourceModel] `tfsdk:"data" json:"data,computed"`
}

type TokenDataDataSourceModel struct {
	ID      types.String `tfsdk:"id" json:"id,computed"`
	Token   types.String `tfsdk:"token" json:"token,computed"`
	Created types.String `tfsdk:"created" json:"created,computed"`
}
