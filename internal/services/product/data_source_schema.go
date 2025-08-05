// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package product

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/terminaldotshop/terraform-provider-terminal/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*ProductDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "ID of the product to get.",
				Required:    true,
			},
			"data": schema.SingleNestedAttribute{
				Description: "Product sold in the Terminal shop.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[ProductDataDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description: "Unique object identifier.\nThe format and length of IDs may change over time.",
						Computed:    true,
					},
					"description": schema.StringAttribute{
						Description: "Description of the product.",
						Computed:    true,
					},
					"name": schema.StringAttribute{
						Description: "Name of the product.",
						Computed:    true,
					},
					"variants": schema.ListNestedAttribute{
						Description: "List of variants of the product.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectListType[ProductDataVariantsDataSourceModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "Unique object identifier.\nThe format and length of IDs may change over time.",
									Computed:    true,
								},
								"name": schema.StringAttribute{
									Description: "Name of the product variant.",
									Computed:    true,
								},
								"price": schema.Int64Attribute{
									Description: "Price of the product variant in cents (USD).",
									Computed:    true,
									Validators: []validator.Int64{
										int64validator.AtLeast(0),
									},
								},
								"tags": schema.SingleNestedAttribute{
									Description: "Tags for the product variant.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[ProductDataVariantsTagsDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"app": schema.StringAttribute{
											Computed: true,
										},
										"market_eu": schema.BoolAttribute{
											Computed: true,
										},
										"market_global": schema.BoolAttribute{
											Computed: true,
										},
										"market_na": schema.BoolAttribute{
											Computed: true,
										},
									},
								},
							},
						},
					},
					"order": schema.Int64Attribute{
						Description: "Order of the product used when displaying a sorted list of products.",
						Computed:    true,
					},
					"subscription": schema.StringAttribute{
						Description: "Whether the product must be or can be subscribed to.\nAvailable values: \"allowed\", \"required\".",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("allowed", "required"),
						},
					},
					"tags": schema.SingleNestedAttribute{
						Description: "Tags for the product.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[ProductDataTagsDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"app": schema.StringAttribute{
								Computed: true,
							},
							"color": schema.StringAttribute{
								Computed: true,
							},
							"featured": schema.BoolAttribute{
								Computed: true,
							},
							"market_eu": schema.BoolAttribute{
								Computed: true,
							},
							"market_global": schema.BoolAttribute{
								Computed: true,
							},
							"market_na": schema.BoolAttribute{
								Computed: true,
							},
						},
					},
					"time_hidden": schema.StringAttribute{
						Description: "Timestamp when the product was hidden from public view.",
						Computed:    true,
					},
				},
			},
		},
	}
}

func (d *ProductDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ProductDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
