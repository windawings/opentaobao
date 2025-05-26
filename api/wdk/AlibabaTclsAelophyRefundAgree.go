package wdk

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/wdk"
)

// AlibabaTclsAelophyRefundAgree saas 售后逆向 商户同意用户逆向申请
// alibaba.tcls.aelophy.refund.agree
//
// saas 售后逆向 商户同意用户逆向申请
func AlibabaTclsAelophyRefundAgree(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaTclsAelophyRefundAgreeAPIRequest, resp *wdk.AlibabaTclsAelophyRefundAgreeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
