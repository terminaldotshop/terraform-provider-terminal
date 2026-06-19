// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package subscription

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/terminaldotshop/terraform-provider-terminal/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*SubscriptionDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "ID of the subscription to get.",
				Required:    true,
			},
			"data": schema.SingleNestedAttribute{
				Description: "Subscription to a Terminal shop product.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[SubscriptionDataDataSourceModel](ctx),
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
						CustomType:  customfield.NewNestedObjectType[SubscriptionDataScheduleDataSourceModel](ctx),
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

func (d *SubscriptionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *SubscriptionDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
