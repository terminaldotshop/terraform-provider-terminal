// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package order

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/terminaldotshop/terraform-provider-terminal/internal/customfield"
)

var _ resource.ResourceWithConfigValidators = (*OrderResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "ID of the order to get.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"address_id": schema.StringAttribute{
				Description:   "Shipping address ID.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"card_id": schema.StringAttribute{
				Description:   "Card ID.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"variants": schema.MapAttribute{
				Description:   "Product variants to include in the order, along with their quantities.",
				Required:      true,
				ElementType:   types.Int64Type,
				PlanModifiers: []planmodifier.Map{mapplanmodifier.RequiresReplace()},
			},
			"data": schema.SingleNestedAttribute{
				Description: "An order from the Terminal shop.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[OrderDataModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description: "Unique object identifier.\nThe format and length of IDs may change over time.",
						Computed:    true,
					},
					"amount": schema.SingleNestedAttribute{
						Description: "The subtotal and shipping amounts of the order.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[OrderDataAmountModel](ctx),
						Attributes: map[string]schema.Attribute{
							"shipping": schema.Int64Attribute{
								Description: "Shipping amount of the order, in cents (USD).",
								Computed:    true,
							},
							"subtotal": schema.Int64Attribute{
								Description: "Subtotal amount of the order, in cents (USD).",
								Computed:    true,
							},
						},
					},
					"items": schema.ListNestedAttribute{
						Description: "Items in the order.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectListType[OrderDataItemsModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "Unique object identifier.\nThe format and length of IDs may change over time.",
									Computed:    true,
								},
								"amount": schema.Int64Attribute{
									Description: "Amount of the item in the order, in cents (USD).",
									Computed:    true,
								},
								"quantity": schema.Int64Attribute{
									Description: "Quantity of the item in the order.",
									Computed:    true,
									Validators: []validator.Int64{
										int64validator.AtLeast(0),
									},
								},
								"description": schema.StringAttribute{
									Description: "Description of the item in the order.",
									Computed:    true,
								},
								"product_variant_id": schema.StringAttribute{
									Description: "ID of the product variant of the item in the order.",
									Computed:    true,
								},
							},
						},
					},
					"shipping": schema.SingleNestedAttribute{
						Description: "Shipping address of the order.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[OrderDataShippingModel](ctx),
						Attributes: map[string]schema.Attribute{
							"city": schema.StringAttribute{
								Description: "City of the address.",
								Computed:    true,
							},
							"country": schema.StringAttribute{
								Description: "ISO 3166-1 alpha-2 country code of the address.",
								Computed:    true,
							},
							"name": schema.StringAttribute{
								Description: "The recipient's name.",
								Computed:    true,
							},
							"street1": schema.StringAttribute{
								Description: "Street of the address.",
								Computed:    true,
							},
							"zip": schema.StringAttribute{
								Description: "Zip code of the address.",
								Computed:    true,
							},
							"phone": schema.StringAttribute{
								Description: "Phone number of the recipient.",
								Computed:    true,
							},
							"province": schema.StringAttribute{
								Description: "Province or state of the address.",
								Computed:    true,
							},
							"street2": schema.StringAttribute{
								Description: "Apartment, suite, etc. of the address.",
								Computed:    true,
							},
						},
					},
					"tracking": schema.SingleNestedAttribute{
						Description: "Tracking information of the order.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[OrderDataTrackingModel](ctx),
						Attributes: map[string]schema.Attribute{
							"number": schema.StringAttribute{
								Description: "Tracking number of the order.",
								Computed:    true,
							},
							"service": schema.StringAttribute{
								Description: "Shipping service of the order.",
								Computed:    true,
							},
							"url": schema.StringAttribute{
								Description: "Tracking URL of the order.",
								Computed:    true,
							},
						},
					},
					"index": schema.Int64Attribute{
						Description: "Zero-based index of the order for this user only.",
						Computed:    true,
					},
				},
			},
		},
	}
}

func (r *OrderResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *OrderResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
