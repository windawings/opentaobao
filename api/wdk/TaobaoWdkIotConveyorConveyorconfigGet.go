package wdk

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/wdk"
)

// TaobaoWdkIotConveyorConveyorconfigGet 获取悬挂链基本配置信息
// taobao.wdk.iot.conveyor.conveyorconfig.get
//
// 用于从云端WCS获取悬挂链基本配置信息
func TaobaoWdkIotConveyorConveyorconfigGet(ctx context.Context, clt *core.SDKClient, req *wdk.TaobaoWdkIotConveyorConveyorconfigGetAPIRequest, resp *wdk.TaobaoWdkIotConveyorConveyorconfigGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
