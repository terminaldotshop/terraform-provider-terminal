// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package address_test

import (
	"context"
	"testing"

	"github.com/terminaldotshop/terraform-provider-terminal/internal/services/address"
	"github.com/terminaldotshop/terraform-provider-terminal/internal/test_helpers"
)

func TestAddressDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*address.AddressDataSourceModel)(nil)
	schema := address.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
