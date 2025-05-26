package tvupadmin

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tvupadmin"
)

// YunosTvpubadminDeviceApkinfo 获取停开服apk信息
// yunos.tvpubadmin.device.apkinfo
//
// 获取停开服apk信息
func YunosTvpubadminDeviceApkinfo(ctx context.Context, clt *core.SDKClient, req *tvupadmin.YunosTvpubadminDeviceApkinfoAPIRequest, resp *tvupadmin.YunosTvpubadminDeviceApkinfoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
