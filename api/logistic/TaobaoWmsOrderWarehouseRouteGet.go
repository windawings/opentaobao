package logistic

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/logistic"
)

// TaobaoWmsOrderWarehouseRouteGet 获取订单仓库路由信息
// taobao.wms.order.warehouse.route.get
//
// 获取订单仓库路由信息
func TaobaoWmsOrderWarehouseRouteGet(ctx context.Context, clt *core.SDKClient, req *logistic.TaobaoWmsOrderWarehouseRouteGetAPIRequest, resp *logistic.TaobaoWmsOrderWarehouseRouteGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
