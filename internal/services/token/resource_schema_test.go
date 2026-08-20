// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package token_test

import (
	"context"
	"testing"

	"github.com/terminaldotshop/terraform-provider-terminal/internal/services/token"
	"github.com/terminaldotshop/terraform-provider-terminal/internal/test_helpers"
)

func TestTokenModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*token.TokenModel)(nil)
	schema := token.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
