package wdk

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/wdk"
)

// AlibabaWdkElemeBillGet 饿了么日维度对账单查询
// alibaba.wdk.eleme.bill.get
//
// 查询饿了么日维度对账单信息
func AlibabaWdkElemeBillGet(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkElemeBillGetAPIRequest, resp *wdk.AlibabaWdkElemeBillGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
