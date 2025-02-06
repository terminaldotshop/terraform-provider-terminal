// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_test

import (
	"context"
	"testing"

	"github.com/terminaldotshop/terraform-provider-terminal/internal/services/email"
	"github.com/terminaldotshop/terraform-provider-terminal/internal/test_helpers"
)

func TestEmailModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*email.EmailModel)(nil)
	schema := email.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
