// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package card

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var _ resource.ResourceWithConfigValidators = (*CardResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Stripe card token. Learn how to [create one here](https://docs.stripe.com/api/tokens/create_card).",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"token": schema.StringAttribute{
				Description:   "Stripe card token. Learn how to [create one here](https://docs.stripe.com/api/tokens/create_card).",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"data": schema.StringAttribute{
				Description: "ID of the card.",
				Computed:    true,
			},
		},
	}
}

func (r *CardResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *CardResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
