// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package app_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/terminal-terraform/internal/services/app"
	"github.com/stainless-sdks/terminal-terraform/internal/test_helpers"
)

func TestAppDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*app.AppDataSourceModel)(nil)
	schema := app.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
