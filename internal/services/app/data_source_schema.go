// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package app

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/stainless-sdks/terminal-terraform/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*AppDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "ID of the app to get.",
				Required:    true,
			},
			"data": schema.SingleNestedAttribute{
				Description: "A Terminal App used for configuring an OAuth 2.0 client.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[AppDataDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description: "Unique object identifier.\nThe format and length of IDs may change over time.",
						Computed:    true,
					},
					"name": schema.StringAttribute{
						Description: "Name of the app.",
						Computed:    true,
					},
					"redirect_uri": schema.StringAttribute{
						Description: "Redirect URI of the app.",
						Computed:    true,
					},
				},
			},
		},
	}
}

func (d *AppDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *AppDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
