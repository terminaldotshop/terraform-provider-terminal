// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package order

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*OrderResource)(nil)

func (r *OrderResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{}
}
