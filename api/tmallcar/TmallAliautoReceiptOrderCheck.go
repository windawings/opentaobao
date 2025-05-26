package tmallcar

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tmallcar"
)

// TmallAliautoReceiptOrderCheck 查看工单查询订单是否已付款
// tmall.aliauto.receipt.order.check
//
// 查看工单查询订单是否已付款
func TmallAliautoReceiptOrderCheck(ctx context.Context, clt *core.SDKClient, req *tmallcar.TmallAliautoReceiptOrderCheckAPIRequest, resp *tmallcar.TmallAliautoReceiptOrderCheckAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
