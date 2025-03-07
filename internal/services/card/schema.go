// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package card

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/terminaldotshop/terraform-provider-terminal/internal/customfield"
)

var _ resource.ResourceWithConfigValidators = (*CardResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Stripe card token. Learn how to [create one here](https://docs.stripe.com/api/tokens/create_card).",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"token": schema.StringAttribute{
				Description:   "Stripe card token. Learn how to [create one here](https://docs.stripe.com/api/tokens/create_card).",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"data": schema.SingleNestedAttribute{
				Description: "Credit card used for payments in the Terminal shop.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[CardDataModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description: "Unique object identifier.\nThe format and length of IDs may change over time.",
						Computed:    true,
					},
					"brand": schema.StringAttribute{
						Description: "Brand of the card.",
						Computed:    true,
					},
					"expiration": schema.SingleNestedAttribute{
						Description: "Expiration of the card.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[CardDataExpirationModel](ctx),
						Attributes: map[string]schema.Attribute{
							"month": schema.Int64Attribute{
								Description: "Expiration month of the card.",
								Computed:    true,
							},
							"year": schema.Int64Attribute{
								Description: "Expiration year of the card.",
								Computed:    true,
							},
						},
					},
					"last4": schema.StringAttribute{
						Description: "Last four digits of the card.",
						Computed:    true,
					},
				},
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
