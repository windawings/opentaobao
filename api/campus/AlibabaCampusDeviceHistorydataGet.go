package campus

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/campus"
)

// AlibabaCampusDeviceHistorydataGet 设备历史数据批量获取
// alibaba.campus.device.historydata.get
//
// 设备历史数据批量获取
func AlibabaCampusDeviceHistorydataGet(ctx context.Context, clt *core.SDKClient, req *campus.AlibabaCampusDeviceHistorydataGetAPIRequest, resp *campus.AlibabaCampusDeviceHistorydataGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
