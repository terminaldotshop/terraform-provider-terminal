// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package app_test

import (
	"context"
	"testing"

	"github.com/terminaldotshop/terraform-provider-terminal/internal/services/app"
	"github.com/terminaldotshop/terraform-provider-terminal/internal/test_helpers"
)

func TestAppModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*app.AppModel)(nil)
	schema := app.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
