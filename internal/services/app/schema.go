// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package app

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/terminaldotshop/terraform-provider-terminal/internal/customfield"
)

var _ resource.ResourceWithConfigValidators = (*AppResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "ID of the app to get.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"redirect_uri": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"data": schema.SingleNestedAttribute{
				Description: "A Terminal App used for configuring an OAuth 2.0 client.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[AppDataModel](ctx),
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
					"secret": schema.StringAttribute{
						Description: "OAuth 2.0 client secret of the app (obfuscated).",
						Computed:    true,
					},
				},
			},
		},
	}
}

func (r *AppResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *AppResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
