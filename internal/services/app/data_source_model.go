// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package app

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/terminal-terraform/internal/customfield"
)

type AppDataSourceModel struct {
	ID   types.String                                     `tfsdk:"id" path:"id,required"`
	Data customfield.NestedObject[AppDataDataSourceModel] `tfsdk:"data" json:"data,computed"`
}

type AppDataDataSourceModel struct {
	ID          types.String `tfsdk:"id" json:"id,computed"`
	Name        types.String `tfsdk:"name" json:"name,computed"`
	RedirectUri types.String `tfsdk:"redirect_uri" json:"redirectURI,computed"`
}
