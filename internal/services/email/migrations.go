// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*EmailResource)(nil)

func (r *EmailResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{}
}
