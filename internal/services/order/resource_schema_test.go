// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package order_test

import (
	"context"
	"testing"

	"github.com/terminaldotshop/terraform-provider-terminal/internal/services/order"
	"github.com/terminaldotshop/terraform-provider-terminal/internal/test_helpers"
)

func TestOrderModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*order.OrderModel)(nil)
	schema := order.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
