// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package address

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/terminaldotshop/terraform-provider-terminal/internal/customfield"
)

var _ resource.ResourceWithConfigValidators = (*AddressResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "ID of the shipping address to get.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"city": schema.StringAttribute{
				Description:   "City of the address.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"country": schema.StringAttribute{
				Description:   "ISO 3166-1 alpha-2 country code of the address.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Description:   "The recipient's name.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"street1": schema.StringAttribute{
				Description:   "Street of the address.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"zip": schema.StringAttribute{
				Description:   "Zip code of the address.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"phone": schema.StringAttribute{
				Description:   "Phone number of the recipient.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"province": schema.StringAttribute{
				Description:   "Province or state of the address.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"street2": schema.StringAttribute{
				Description:   "Apartment, suite, etc. of the address.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"data": schema.SingleNestedAttribute{
				Description: "Physical address associated with a Terminal shop user.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[AddressDataModel](ctx),
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

func (r *AddressResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *AddressResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
