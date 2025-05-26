package wdk

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/wdk"
)

// WdkWarehouseOrderDispatch 仓作业下发
// wdk.warehouse.order.dispatch
//
// 牵牛花仓作业下发接口提供
func WdkWarehouseOrderDispatch(ctx context.Context, clt *core.SDKClient, req *wdk.WdkWarehouseOrderDispatchAPIRequest, resp *wdk.WdkWarehouseOrderDispatchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
