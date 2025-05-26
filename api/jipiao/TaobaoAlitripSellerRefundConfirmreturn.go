package jipiao

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/jipiao"
)

// TaobaoAlitripSellerRefundConfirmreturn 【机票代理商】确认退票
// taobao.alitrip.seller.refund.confirmreturn
//
// 确认退票
func TaobaoAlitripSellerRefundConfirmreturn(ctx context.Context, clt *core.SDKClient, req *jipiao.TaobaoAlitripSellerRefundConfirmreturnAPIRequest, resp *jipiao.TaobaoAlitripSellerRefundConfirmreturnAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
