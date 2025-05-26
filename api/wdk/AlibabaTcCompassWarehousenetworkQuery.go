package wdk

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/wdk"
)

// AlibabaTcCompassWarehousenetworkQuery 按仓维度来查询鸟潮网络
// alibaba.tc.compass.warehousenetwork.query
//
// 按仓维度来查询鸟潮网络
func AlibabaTcCompassWarehousenetworkQuery(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaTcCompassWarehousenetworkQueryAPIRequest, resp *wdk.AlibabaTcCompassWarehousenetworkQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
