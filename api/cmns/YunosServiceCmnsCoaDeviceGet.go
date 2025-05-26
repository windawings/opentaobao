package cmns

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/cmns"
)

// YunosServiceCmnsCoaDeviceGet 设备详情查询
// yunos.service.cmns.coa.device.get
//
// 第三方应用开发者调用此接口查询设备详情
func YunosServiceCmnsCoaDeviceGet(ctx context.Context, clt *core.SDKClient, req *cmns.YunosServiceCmnsCoaDeviceGetAPIRequest, resp *cmns.YunosServiceCmnsCoaDeviceGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
