// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package order

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/terminaldotshop/terraform-provider-terminal/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*OrderDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "ID of the order to get.",
				Required:    true,
			},
			"data": schema.SingleNestedAttribute{
				Description: "An order from the Terminal shop.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[OrderDataDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description: "Unique object identifier.\nThe format and length of IDs may change over time.",
						Computed:    true,
					},
					"amount": schema.SingleNestedAttribute{
						Description: "The subtotal and shipping amounts of the order.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[OrderDataAmountDataSourceModel](ctx),
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
						CustomType:  customfield.NewNestedObjectListType[OrderDataItemsDataSourceModel](ctx),
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
						CustomType:  customfield.NewNestedObjectType[OrderDataShippingDataSourceModel](ctx),
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
						CustomType:  customfield.NewNestedObjectType[OrderDataTrackingDataSourceModel](ctx),
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

func (d *OrderDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *OrderDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
