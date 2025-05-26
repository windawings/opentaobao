package icbudropshipping

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/icbudropshipping"
)

// AlibabaOrderPayResultQuery alibaba查询订单支付结果
// alibaba.order.pay.result.query
//
// alibaba查询订单支付结果
func AlibabaOrderPayResultQuery(ctx context.Context, clt *core.SDKClient, req *icbudropshipping.AlibabaOrderPayResultQueryAPIRequest, resp *icbudropshipping.AlibabaOrderPayResultQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
