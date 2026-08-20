// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package token

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/terminaldotshop/terraform-provider-terminal/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*TokenDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "ID of the personal token to get.",
				Required:    true,
			},
			"data": schema.SingleNestedAttribute{
				Description: "A personal access token used to access the Terminal API. If you leak this, expect large sums of coffee to be ordered on your credit card.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[TokenDataDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description: "Unique object identifier.\nThe format and length of IDs may change over time.",
						Computed:    true,
					},
					"token": schema.StringAttribute{
						Description: "Personal access token (obfuscated).",
						Computed:    true,
					},
					"created": schema.StringAttribute{
						Description: "The created time for the token.",
						Computed:    true,
					},
				},
			},
		},
	}
}

func (d *TokenDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *TokenDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
