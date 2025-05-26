package tbrefund

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tbrefund"
)

// TaobaoRpRefundIntercept 卖家发起拦截
// taobao.rp.refund.intercept
//
// 卖家发起拦截
func TaobaoRpRefundIntercept(ctx context.Context, clt *core.SDKClient, req *tbrefund.TaobaoRpRefundInterceptAPIRequest, resp *tbrefund.TaobaoRpRefundInterceptAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
