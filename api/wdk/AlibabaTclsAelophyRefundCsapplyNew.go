package wdk

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/wdk"
)

// AlibabaTclsAelophyRefundCsapplyNew 代客退
// alibaba.tcls.aelophy.refund.csapply.new
//
// 代客退
func AlibabaTclsAelophyRefundCsapplyNew(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaTclsAelophyRefundCsapplyNewAPIRequest, resp *wdk.AlibabaTclsAelophyRefundCsapplyNewAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
