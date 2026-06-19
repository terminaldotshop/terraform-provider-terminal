// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package address

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/terminaldotshop/terraform-provider-terminal/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*AddressDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "ID of the shipping address to get.",
				Required:    true,
			},
			"data": schema.SingleNestedAttribute{
				Description: "Physical address associated with a Terminal shop user.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[AddressDataDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description: "Unique object identifier.\nThe format and length of IDs may change over time.",
						Computed:    true,
					},
					"city": schema.StringAttribute{
						Description: "City of the address.",
						Computed:    true,
					},
					"country": schema.StringAttribute{
						Description: "ISO 3166-1 alpha-2 country code of the address.",
						Computed:    true,
					},
					"created": schema.StringAttribute{
						Description: "Date the address was created.",
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
		},
	}
}

func (d *AddressDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *AddressDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
