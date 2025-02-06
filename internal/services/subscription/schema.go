// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package subscription

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*SubscriptionResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Unique object identifier.\nThe format and length of IDs may change over time.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"address_id": schema.StringAttribute{
				Description: "ID of the shipping address used for the subscription.",
				Required:    true,
			},
			"card_id": schema.StringAttribute{
				Description: "ID of the card used for the subscription.",
				Required:    true,
			},
			"frequency": schema.StringAttribute{
				Description: "Frequency of the subscription.",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"fixed",
						"daily",
						"weekly",
						"monthly",
						"yearly",
					),
				},
			},
			"product_variant_id": schema.StringAttribute{
				Description: "ID of the product variant being subscribed to.",
				Required:    true,
			},
			"quantity": schema.Int64Attribute{
				Description: "Quantity of the subscription.",
				Required:    true,
			},
			"next": schema.StringAttribute{
				Description: "Next shipment and billing date for the subscription.",
				Optional:    true,
			},
			"data": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("ok"),
				},
			},
		},
	}
}

func (r *SubscriptionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *SubscriptionResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
