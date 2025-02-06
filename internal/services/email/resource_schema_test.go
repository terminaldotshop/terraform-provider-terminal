// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/terminal-terraform/internal/services/email"
	"github.com/stainless-sdks/terminal-terraform/internal/test_helpers"
)

func TestEmailModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*email.EmailModel)(nil)
	schema := email.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
