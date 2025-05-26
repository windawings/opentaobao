package tvupadmin

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tvupadmin"
)

// YunosTvpubadminDeviceApks 获取停开服apk列表
// yunos.tvpubadmin.device.apks
//
// 获取停开服apk列表
func YunosTvpubadminDeviceApks(ctx context.Context, clt *core.SDKClient, req *tvupadmin.YunosTvpubadminDeviceApksAPIRequest, resp *tvupadmin.YunosTvpubadminDeviceApksAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
