// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package card

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/terminaldotshop/terraform-provider-terminal/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*CardDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "ID of the card to get.",
				Required:    true,
			},
			"data": schema.SingleNestedAttribute{
				Description: "Credit card used for payments in the Terminal shop.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[CardDataDataSourceModel](ctx),
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
						CustomType:  customfield.NewNestedObjectType[CardDataExpirationDataSourceModel](ctx),
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

func (d *CardDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *CardDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
