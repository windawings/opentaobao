package wdk

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/wdk"
)

// AlibabaWdkWholesaleOrderCommit 盒马帮采购确认订单接口
// alibaba.wdk.wholesale.order.commit
//
// 盒马帮采购确认订单接口
func AlibabaWdkWholesaleOrderCommit(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkWholesaleOrderCommitAPIRequest, resp *wdk.AlibabaWdkWholesaleOrderCommitAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
