package ieagency

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/ieagency"
)

// AlitripIeBuyerOrderBookpay 【国际机票】下单预定支付
// alitrip.ie.buyer.order.bookpay
//
// 【国际机票】 生单预定支付接口
func AlitripIeBuyerOrderBookpay(ctx context.Context, clt *core.SDKClient, req *ieagency.AlitripIeBuyerOrderBookpayAPIRequest, resp *ieagency.AlitripIeBuyerOrderBookpayAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
