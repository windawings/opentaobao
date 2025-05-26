package idle

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/idle"
)

// TaobaoIdleRecycleRefundCancleapply 闲鱼回收取消退款申请V2
// taobao.idle.recycle.refund.cancleapply
//
// 回收商的回收订单取消退款申请
func TaobaoIdleRecycleRefundCancleapply(ctx context.Context, clt *core.SDKClient, req *idle.TaobaoIdleRecycleRefundCancleapplyAPIRequest, resp *idle.TaobaoIdleRecycleRefundCancleapplyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
