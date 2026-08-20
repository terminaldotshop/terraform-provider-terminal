// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package token_test

import (
	"context"
	"testing"

	"github.com/terminaldotshop/terraform-provider-terminal/internal/services/token"
	"github.com/terminaldotshop/terraform-provider-terminal/internal/test_helpers"
)

func TestTokenDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*token.TokenDataSourceModel)(nil)
	schema := token.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
