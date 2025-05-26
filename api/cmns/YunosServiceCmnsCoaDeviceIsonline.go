package cmns

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/cmns"
)

// YunosServiceCmnsCoaDeviceIsonline 根据设备id查询设备是否在线
// yunos.service.cmns.coa.device.isonline
//
// 根据设备id查询设备是否在线
func YunosServiceCmnsCoaDeviceIsonline(ctx context.Context, clt *core.SDKClient, req *cmns.YunosServiceCmnsCoaDeviceIsonlineAPIRequest, resp *cmns.YunosServiceCmnsCoaDeviceIsonlineAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
