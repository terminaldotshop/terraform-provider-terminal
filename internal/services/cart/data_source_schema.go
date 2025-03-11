// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cart

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/terminaldotshop/terraform-provider-terminal/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*CartDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"data": schema.SingleNestedAttribute{
				Description: "The current Terminal shop user's cart.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[CartDataDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"amount": schema.SingleNestedAttribute{
						Description: "The subtotal and shipping amounts for the current user's cart.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[CartDataAmountDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"subtotal": schema.Int64Attribute{
								Description: "Subtotal of the current user's cart, in cents (USD).",
								Computed:    true,
							},
							"gift_card": schema.Int64Attribute{
								Description: "Amount applied from gift card on the current user's cart, in cents (USD).",
								Computed:    true,
							},
							"shipping": schema.Int64Attribute{
								Description: "Shipping amount of the current user's cart, in cents (USD).",
								Computed:    true,
							},
							"total": schema.Int64Attribute{
								Description: "Total amount after gift card applied, in cents (USD).",
								Computed:    true,
							},
						},
					},
					"items": schema.ListNestedAttribute{
						Description: "An array of items in the current user's cart.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectListType[CartDataItemsDataSourceModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "Unique object identifier.\nThe format and length of IDs may change over time.",
									Computed:    true,
								},
								"product_variant_id": schema.StringAttribute{
									Description: "ID of the product variant for this item in the current user's cart.",
									Computed:    true,
								},
								"quantity": schema.Int64Attribute{
									Description: "Quantity of the item in the current user's cart.",
									Computed:    true,
									Validators: []validator.Int64{
										int64validator.AtLeast(0),
									},
								},
								"subtotal": schema.Int64Attribute{
									Description: "Subtotal of the item in the current user's cart, in cents (USD).",
									Computed:    true,
								},
							},
						},
					},
					"subtotal": schema.Int64Attribute{
						Description: "The subtotal of all items in the current user's cart, in cents (USD).",
						Computed:    true,
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
						},
					},
					"address_id": schema.StringAttribute{
						Description: "ID of the shipping address selected on the current user's cart.",
						Computed:    true,
					},
					"card_id": schema.StringAttribute{
						Description: "ID of the card selected on the current user's cart.",
						Computed:    true,
					},
					"gift_card_id": schema.StringAttribute{
						Description: "ID of the gift card applied to the current user's cart.",
						Computed:    true,
					},
					"shipping": schema.SingleNestedAttribute{
						Description: "Shipping information for the current user's cart.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[CartDataShippingDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"service": schema.StringAttribute{
								Description: "Shipping service name.",
								Computed:    true,
							},
							"timeframe": schema.StringAttribute{
								Description: "Shipping timeframe provided by the shipping carrier.",
								Computed:    true,
							},
						},
					},
				},
			},
		},
	}
}

func (d *CartDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *CartDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
