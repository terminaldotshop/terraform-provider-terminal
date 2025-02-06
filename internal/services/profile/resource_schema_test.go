// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package profile_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/terminal-terraform/internal/services/profile"
	"github.com/stainless-sdks/terminal-terraform/internal/test_helpers"
)

func TestProfileModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*profile.ProfileModel)(nil)
	schema := profile.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
