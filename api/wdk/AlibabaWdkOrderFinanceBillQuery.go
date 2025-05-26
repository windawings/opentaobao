package wdk

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/wdk"
)

// AlibabaWdkOrderFinanceBillQuery 资金合规商家账单
// alibaba.wdk.order.finance.bill.query
//
// 拉取资金合规商家账单
func AlibabaWdkOrderFinanceBillQuery(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkOrderFinanceBillQueryAPIRequest, resp *wdk.AlibabaWdkOrderFinanceBillQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
