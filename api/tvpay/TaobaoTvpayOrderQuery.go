package tvpay

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tvpay"
)

// TaobaoTvpayOrderQuery tv支付查询订单状态
// taobao.tvpay.order.query
//
// tv支付查询订单状态
func TaobaoTvpayOrderQuery(ctx context.Context, clt *core.SDKClient, req *tvpay.TaobaoTvpayOrderQueryAPIRequest, resp *tvpay.TaobaoTvpayOrderQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
