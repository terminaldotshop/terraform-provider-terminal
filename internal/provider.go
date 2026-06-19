// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package internal

import (
	"context"
	"fmt"
	"os"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/terminaldotshop/terminal-sdk-go"
	"github.com/terminaldotshop/terminal-sdk-go/option"
	"github.com/terminaldotshop/terraform-provider-terminal/internal/services/address"
	"github.com/terminaldotshop/terraform-provider-terminal/internal/services/app"
	"github.com/terminaldotshop/terraform-provider-terminal/internal/services/card"
	"github.com/terminaldotshop/terraform-provider-terminal/internal/services/cart"
	"github.com/terminaldotshop/terraform-provider-terminal/internal/services/email"
	"github.com/terminaldotshop/terraform-provider-terminal/internal/services/order"
	"github.com/terminaldotshop/terraform-provider-terminal/internal/services/product"
	"github.com/terminaldotshop/terraform-provider-terminal/internal/services/profile"
	"github.com/terminaldotshop/terraform-provider-terminal/internal/services/subscription"
	"github.com/terminaldotshop/terraform-provider-terminal/internal/services/token"
)

var _ provider.ProviderWithConfigValidators = (*TerminalProvider)(nil)

// TerminalProvider defines the provider implementation.
type TerminalProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// TerminalProviderModel describes the provider data model.
type TerminalProviderModel struct {
	BaseURL     types.String `tfsdk:"base_url" json:"base_url,optional"`
	Environment types.String `tfsdk:"environment" json:"environment,optional"`
	BearerToken types.String `tfsdk:"bearer_token" json:"bearer_token,optional"`
	AppID       types.String `tfsdk:"app_id" json:"app_id,optional"`
}

func (p *TerminalProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "terminal"
	resp.Version = p.version
}

func ProviderSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"base_url": schema.StringAttribute{
				Description: "Set the base url that the provider connects to.",
				Optional:    true,
			},
			"environment": schema.StringAttribute{
				Description: "Set the current environment. An environment specifies which base URL to use by default.\nAvailable values: \"production\", \"dev\".",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("production", "dev"),
				},
			},
			"bearer_token": schema.StringAttribute{
				Optional: true,
			},
			"app_id": schema.StringAttribute{
				Optional: true,
			},
		},
	}
}

func (p *TerminalProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = ProviderSchema(ctx)
}

func (p *TerminalProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {

	var data TerminalProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	opts := []option.RequestOption{}

	if !data.BaseURL.IsNull() && !data.BaseURL.IsUnknown() {
		opts = append(opts, option.WithBaseURL(data.BaseURL.ValueString()))
	} else if o, ok := os.LookupEnv("TERMINAL_BASE_URL"); ok {
		opts = append(opts, option.WithBaseURL(o))
	}

	if !data.Environment.IsNull() && !data.Environment.IsUnknown() {
		environment := data.Environment.ValueString()

		// First check for conflict between environment and base_url settings
		if !data.BaseURL.IsNull() && !data.BaseURL.IsUnknown() {
			resp.Diagnostics.AddAttributeError(path.Root("environment"),
				"Ambiguous URL: the 'base_url' and 'environment' options are both set in the provider configuration.",
				"At most one of these fields may be set.",
			)
		} else if o, ok := os.LookupEnv("TERMINAL_BASE_URL"); ok {
			resp.Diagnostics.AddAttributeWarning(path.Root("environment"),
				"Base URL from TERMINAL_BASE_URL environment variable overrides 'environment' option from provider configuration.",
				fmt.Sprintf("Using base URL %s", o),
			)
		}

		// Then add client option for specified environment
		switch environment {
		case "production":
			opts = append(opts, option.WithEnvironmentProduction())
		case "dev":
			opts = append(opts, option.WithEnvironmentDev())
		}
	}

	if !data.BearerToken.IsNull() && !data.BearerToken.IsUnknown() {
		opts = append(opts, option.WithBearerToken(data.BearerToken.ValueString()))
	} else if o, ok := os.LookupEnv("TERMINAL_BEARER_TOKEN"); ok {
		opts = append(opts, option.WithBearerToken(o))
	} else {
		resp.Diagnostics.AddAttributeError(
			path.Root("bearer_token"),
			"Missing bearer_token value",
			"The bearer_token field is required. Set it in provider configuration or via the \"TERMINAL_BEARER_TOKEN\" environment variable.",
		)
		return
	}

	if !data.AppID.IsNull() && !data.AppID.IsUnknown() {
		opts = append(opts, option.WithAppID(data.AppID.ValueString()))
	}

	client := githubcomterminaldotshopterminalsdkgo.NewClient(
		opts...,
	)

	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *TerminalProvider) ConfigValidators(_ context.Context) []provider.ConfigValidator {
	return []provider.ConfigValidator{}
}

func (p *TerminalProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		profile.NewResource,
		address.NewResource,
		card.NewResource,
		order.NewResource,
		subscription.NewResource,
		token.NewResource,
		app.NewResource,
		email.NewResource,
	}
}

func (p *TerminalProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		product.NewProductDataSource,
		address.NewAddressDataSource,
		card.NewCardDataSource,
		cart.NewCartDataSource,
		order.NewOrderDataSource,
		subscription.NewSubscriptionDataSource,
		token.NewTokenDataSource,
		app.NewAppDataSource,
	}
}

func NewProvider(version string) func() provider.Provider {
	return func() provider.Provider {
		return &TerminalProvider{
			version: version,
		}
	}
}
