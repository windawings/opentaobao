package wdk

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/wdk"
)

// AlibabaWdkOrderList 五道口订单拉取
// alibaba.wdk.order.list
//
// 五道口交易订单拉取接口
func AlibabaWdkOrderList(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkOrderListAPIRequest, resp *wdk.AlibabaWdkOrderListAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
