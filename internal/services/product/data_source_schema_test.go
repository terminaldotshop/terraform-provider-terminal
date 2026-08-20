// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package product_test

import (
	"context"
	"testing"

	"github.com/terminaldotshop/terraform-provider-terminal/internal/services/product"
	"github.com/terminaldotshop/terraform-provider-terminal/internal/test_helpers"
)

func TestProductDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*product.ProductDataSourceModel)(nil)
	schema := product.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
