package bus

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/bus"
)

// TaobaoBusRefundticketpriceSet 汽车票退款申请接口
// taobao.bus.refundticketprice.set
//
// 汽车票代理商利用该接口申请退票
func TaobaoBusRefundticketpriceSet(ctx context.Context, clt *core.SDKClient, req *bus.TaobaoBusRefundticketpriceSetAPIRequest, resp *bus.TaobaoBusRefundticketpriceSetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
