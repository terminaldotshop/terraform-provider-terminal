// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package card_test

import (
	"context"
	"testing"

	"github.com/terminaldotshop/terraform-provider-terminal/internal/services/card"
	"github.com/terminaldotshop/terraform-provider-terminal/internal/test_helpers"
)

func TestCardDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*card.CardDataSourceModel)(nil)
	schema := card.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
