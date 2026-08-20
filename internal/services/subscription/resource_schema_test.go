// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package subscription_test

import (
	"context"
	"testing"

	"github.com/terminaldotshop/terraform-provider-terminal/internal/services/subscription"
	"github.com/terminaldotshop/terraform-provider-terminal/internal/test_helpers"
)

func TestSubscriptionModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*subscription.SubscriptionModel)(nil)
	schema := subscription.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
