// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package profile

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/terminaldotshop/terraform-provider-terminal/internal/apijson"
	"github.com/terminaldotshop/terraform-provider-terminal/internal/customfield"
)

type ProfileModel struct {
	Email types.String                               `tfsdk:"email" json:"email,optional"`
	Name  types.String                               `tfsdk:"name" json:"name,optional"`
	Data  customfield.NestedObject[ProfileDataModel] `tfsdk:"data" json:"data,computed"`
}

func (m ProfileModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ProfileModel) MarshalJSONForUpdate(state ProfileModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type ProfileDataModel struct {
	User customfield.NestedObject[ProfileDataUserModel] `tfsdk:"user" json:"user,computed"`
}

type ProfileDataUserModel struct {
	ID               types.String `tfsdk:"id" json:"id,computed"`
	Email            types.String `tfsdk:"email" json:"email,computed"`
	Fingerprint      types.String `tfsdk:"fingerprint" json:"fingerprint,computed"`
	Name             types.String `tfsdk:"name" json:"name,computed"`
	StripeCustomerID types.String `tfsdk:"stripe_customer_id" json:"stripeCustomerID,computed"`
}
