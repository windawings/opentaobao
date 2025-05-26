package tvupadmin

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tvupadmin"
)

// YunosTvpubadminDeviceYksBots 获取设备列表
// yunos.tvpubadmin.device.yks.bots
//
// 获取设备列表
func YunosTvpubadminDeviceYksBots(ctx context.Context, clt *core.SDKClient, req *tvupadmin.YunosTvpubadminDeviceYksBotsAPIRequest, resp *tvupadmin.YunosTvpubadminDeviceYksBotsAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
