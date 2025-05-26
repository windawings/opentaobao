package bill

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/bill"
)

// TaobaoTaeBillGet tae查询单笔账单明细
// taobao.tae.bill.get
//
// 查询单笔账单明细
func TaobaoTaeBillGet(ctx context.Context, clt *core.SDKClient, req *bill.TaobaoTaeBillGetAPIRequest, resp *bill.TaobaoTaeBillGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
