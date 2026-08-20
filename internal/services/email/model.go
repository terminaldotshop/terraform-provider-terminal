// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/terminaldotshop/terraform-provider-terminal/internal/apijson"
)

type EmailModel struct {
	Email types.String `tfsdk:"email" json:"email,required"`
	Data  types.String `tfsdk:"data" json:"data,computed"`
}

func (m EmailModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m EmailModel) MarshalJSONForUpdate(state EmailModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
