// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package token

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/terminaldotshop/terminal-sdk-go"
	"github.com/terminaldotshop/terminal-sdk-go/option"
	"github.com/terminaldotshop/terraform-provider-terminal/internal/apijson"
	"github.com/terminaldotshop/terraform-provider-terminal/internal/logging"
)

type TokenDataSource struct {
	client *terminal.Client
}

var _ datasource.DataSourceWithConfigure = (*TokenDataSource)(nil)

func NewTokenDataSource() datasource.DataSource {
	return &TokenDataSource{}
}

func (d *TokenDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_token"
}

func (d *TokenDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*terminal.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"unexpected resource configure type",
			fmt.Sprintf("Expected *terminal.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}

func (d *TokenDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *TokenDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	res := new(http.Response)
	_, err := d.client.Token.Get(
		ctx,
		data.ID.ValueString(),
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	bytes, _ := io.ReadAll(res.Body)
	err = apijson.UnmarshalComputed(bytes, &data)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
