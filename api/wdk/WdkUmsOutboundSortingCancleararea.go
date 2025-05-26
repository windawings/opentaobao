package wdk

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/wdk"
)

// WdkUmsOutboundSortingCancleararea dps分货-是否能够清场
// wdk.ums.outbound.sorting.cancleararea
//
// dps分货-是否能够清场
func WdkUmsOutboundSortingCancleararea(ctx context.Context, clt *core.SDKClient, req *wdk.WdkUmsOutboundSortingCanclearareaAPIRequest, resp *wdk.WdkUmsOutboundSortingCanclearareaAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
