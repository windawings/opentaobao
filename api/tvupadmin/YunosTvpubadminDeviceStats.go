package tvupadmin

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tvupadmin"
)

// YunosTvpubadminDeviceStats 获取设备统计数据
// yunos.tvpubadmin.device.stats
//
// 获取设备统计数据
func YunosTvpubadminDeviceStats(ctx context.Context, clt *core.SDKClient, req *tvupadmin.YunosTvpubadminDeviceStatsAPIRequest, resp *tvupadmin.YunosTvpubadminDeviceStatsAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
