// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package subscription_test

import (
	"context"
	"testing"

	"github.com/terminaldotshop/terraform-provider-terminal/internal/services/subscription"
	"github.com/terminaldotshop/terraform-provider-terminal/internal/test_helpers"
)

func TestSubscriptionDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*subscription.SubscriptionDataSourceModel)(nil)
	schema := subscription.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
