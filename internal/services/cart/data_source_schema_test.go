// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cart_test

import (
	"context"
	"testing"

	"github.com/terminaldotshop/terraform-provider-terminal/internal/services/cart"
	"github.com/terminaldotshop/terraform-provider-terminal/internal/test_helpers"
)

func TestCartDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*cart.CartDataSourceModel)(nil)
	schema := cart.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
