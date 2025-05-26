package lsttrade

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/lsttrade"
)

// AlibabaLstTradeRefundOrderGet 零售通退款订单查询
// alibaba.lst.trade.refund.order.get
//
// 零售通退款订单查询
func AlibabaLstTradeRefundOrderGet(ctx context.Context, clt *core.SDKClient, req *lsttrade.AlibabaLstTradeRefundOrderGetAPIRequest, resp *lsttrade.AlibabaLstTradeRefundOrderGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
