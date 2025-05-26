package wdk

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/wdk"
)

// AlibabaAelophyOrderDelivererChange 配送员信息变更接口
// alibaba.aelophy.order.deliverer.change
//
// 配送员信息变更接口
func AlibabaAelophyOrderDelivererChange(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaAelophyOrderDelivererChangeAPIRequest, resp *wdk.AlibabaAelophyOrderDelivererChangeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
