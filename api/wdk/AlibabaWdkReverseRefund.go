package wdk

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/wdk"
)

// AlibabaWdkReverseRefund 退款打款
// alibaba.wdk.reverse.refund
//
// 五道口退款打款开放能力接口
func AlibabaWdkReverseRefund(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkReverseRefundAPIRequest, resp *wdk.AlibabaWdkReverseRefundAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
