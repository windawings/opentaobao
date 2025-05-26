package wdk

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/wdk"
)

// AlibabaAelophyOrderGet 翱象拉取订单接口
// alibaba.aelophy.order.get
//
// 翱象拉取订单接口
func AlibabaAelophyOrderGet(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaAelophyOrderGetAPIRequest, resp *wdk.AlibabaAelophyOrderGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
