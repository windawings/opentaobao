package wdk

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/wdk"
)

// AlibabaWdkChannelOrderCreate 创建订单
// alibaba.wdk.channel.order.create
//
// 外部商家创建订单
func AlibabaWdkChannelOrderCreate(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkChannelOrderCreateAPIRequest, resp *wdk.AlibabaWdkChannelOrderCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
