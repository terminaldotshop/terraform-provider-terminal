// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package card

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/terminaldotshop/terraform-provider-terminal/internal/apijson"
	"github.com/terminaldotshop/terraform-provider-terminal/internal/customfield"
)

type CardModel struct {
	ID    types.String                            `tfsdk:"id" json:"-,computed"`
	Token types.String                            `tfsdk:"token" json:"token,required"`
	Data  customfield.NestedObject[CardDataModel] `tfsdk:"data" json:"data,computed"`
}

func (m CardModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m CardModel) MarshalJSONForUpdate(state CardModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type CardDataModel struct {
	ID         types.String                                      `tfsdk:"id" json:"id,computed"`
	Brand      types.String                                      `tfsdk:"brand" json:"brand,computed"`
	Created    types.String                                      `tfsdk:"created" json:"created,computed"`
	Expiration customfield.NestedObject[CardDataExpirationModel] `tfsdk:"expiration" json:"expiration,computed"`
	Last4      types.String                                      `tfsdk:"last4" json:"last4,computed"`
}

type CardDataExpirationModel struct {
	Month types.Int64 `tfsdk:"month" json:"month,computed"`
	Year  types.Int64 `tfsdk:"year" json:"year,computed"`
}
