// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package card

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/terminaldotshop/terraform-provider-terminal/internal/apijson"
)

type CardModel struct {
	ID    types.String `tfsdk:"id" json:"-,computed"`
	Token types.String `tfsdk:"token" json:"token,required"`
	Data  types.String `tfsdk:"data" json:"data,computed"`
}

func (m CardModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m CardModel) MarshalJSONForUpdate(state CardModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
