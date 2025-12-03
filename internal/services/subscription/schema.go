// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package subscription

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/terminaldotshop/terraform-provider-terminal/internal/customfield"
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
			"created": schema.StringAttribute{
				Description:   "Date the subscription was created.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"price": schema.Int64Attribute{
				Description:   "Price of the subscription in cents (USD).",
				Required:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
			"product_variant_id": schema.StringAttribute{
				Description:   "ID of the product variant being subscribed to.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"quantity": schema.Int64Attribute{
				Description:   "Quantity of the subscription.",
				Required:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
			"next": schema.StringAttribute{
				Description:   "Next shipment and billing date for the subscription.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"address_id": schema.StringAttribute{
				Description: "ID of the shipping address used for the subscription.",
				Required:    true,
			},
			"card_id": schema.StringAttribute{
				Description: "ID of the card used for the subscription.",
				Required:    true,
			},
			"schedule": schema.SingleNestedAttribute{
				Description: "Schedule of the subscription.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"type": schema.StringAttribute{
						Description: `Available values: "fixed", "weekly".`,
						Required:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("fixed", "weekly"),
						},
					},
					"interval": schema.Int64Attribute{
						Optional: true,
						Validators: []validator.Int64{
							int64validator.AtLeast(1),
						},
					},
				},
			},
			"data": schema.SingleNestedAttribute{
				Description: "Subscription to a Terminal shop product.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[SubscriptionDataModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description: "Unique object identifier.\nThe format and length of IDs may change over time.",
						Computed:    true,
					},
					"address_id": schema.StringAttribute{
						Description: "ID of the shipping address used for the subscription.",
						Computed:    true,
					},
					"card_id": schema.StringAttribute{
						Description: "ID of the card used for the subscription.",
						Computed:    true,
					},
					"created": schema.StringAttribute{
						Description: "Date the subscription was created.",
						Computed:    true,
					},
					"price": schema.Int64Attribute{
						Description: "Price of the subscription in cents (USD).",
						Computed:    true,
					},
					"product_variant_id": schema.StringAttribute{
						Description: "ID of the product variant being subscribed to.",
						Computed:    true,
					},
					"quantity": schema.Int64Attribute{
						Description: "Quantity of the subscription.",
						Computed:    true,
					},
					"next": schema.StringAttribute{
						Description: "Next shipment and billing date for the subscription.",
						Computed:    true,
					},
					"schedule": schema.SingleNestedAttribute{
						Description: "Schedule of the subscription.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[SubscriptionDataScheduleModel](ctx),
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Description: `Available values: "fixed", "weekly".`,
								Computed:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("fixed", "weekly"),
								},
							},
							"interval": schema.Int64Attribute{
								Computed: true,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
								},
							},
						},
					},
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
