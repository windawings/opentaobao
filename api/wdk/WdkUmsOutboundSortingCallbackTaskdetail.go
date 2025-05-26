package wdk

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/wdk"
)

// WdkUmsOutboundSortingCallbackTaskdetail dps分货，明细回传
// wdk.ums.outbound.sorting.callback.taskdetail
//
// dps分货-分货明细回传
func WdkUmsOutboundSortingCallbackTaskdetail(ctx context.Context, clt *core.SDKClient, req *wdk.WdkUmsOutboundSortingCallbackTaskdetailAPIRequest, resp *wdk.WdkUmsOutboundSortingCallbackTaskdetailAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
