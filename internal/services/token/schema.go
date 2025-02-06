// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package token

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/terminaldotshop/terraform-provider-terminal/internal/customfield"
)

var _ resource.ResourceWithConfigValidators = (*TokenResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "ID of the personal token to get.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"data": schema.SingleNestedAttribute{
				Description: "A personal access token used to access the Terminal API. If you leak this, expect large sums of coffee to be ordered on your credit card.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[TokenDataModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description: "Unique object identifier.\nThe format and length of IDs may change over time.",
						Computed:    true,
					},
					"token": schema.StringAttribute{
						Description: "Personal access token (obfuscated).",
						Computed:    true,
					},
					"time": schema.SingleNestedAttribute{
						Description: "Relevant timestamps for the token.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[TokenDataTimeModel](ctx),
						Attributes: map[string]schema.Attribute{
							"created": schema.StringAttribute{
								Description: "The created time for the token.",
								Computed:    true,
							},
						},
					},
				},
			},
		},
	}
}

func (r *TokenResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *TokenResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
