package wdk

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/wdk"
)

// AlibabaTclsAelophyBillDetailQuery 账单明细接口
// alibaba.tcls.aelophy.bill.detail.query
//
// 账单明细接口
func AlibabaTclsAelophyBillDetailQuery(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaTclsAelophyBillDetailQueryAPIRequest, resp *wdk.AlibabaTclsAelophyBillDetailQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
