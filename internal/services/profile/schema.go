// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package profile

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/terminaldotshop/terraform-provider-terminal/internal/customfield"
)

var _ resource.ResourceWithConfigValidators = (*ProfileResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"email": schema.StringAttribute{
				Description:   "Email address of the user.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Description:   "Name of the user.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"data": schema.SingleNestedAttribute{
				Description: "A Terminal shop user's profile. (We have users, btw.)",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[ProfileDataModel](ctx),
				Attributes: map[string]schema.Attribute{
					"user": schema.SingleNestedAttribute{
						Description: "A Terminal shop user. (We have users, btw.)",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[ProfileDataUserModel](ctx),
						Attributes: map[string]schema.Attribute{
							"id": schema.StringAttribute{
								Description: "Unique object identifier.\nThe format and length of IDs may change over time.",
								Computed:    true,
							},
							"email": schema.StringAttribute{
								Description: "Email address of the user.",
								Computed:    true,
							},
							"fingerprint": schema.StringAttribute{
								Description: "The user's fingerprint, derived from their public SSH key.",
								Computed:    true,
							},
							"name": schema.StringAttribute{
								Description: "Name of the user.",
								Computed:    true,
							},
							"stripe_customer_id": schema.StringAttribute{
								Description: "Stripe customer ID of the user.",
								Computed:    true,
							},
						},
					},
				},
			},
		},
	}
}

func (r *ProfileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ProfileResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
