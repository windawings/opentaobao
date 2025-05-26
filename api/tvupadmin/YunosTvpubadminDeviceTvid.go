package tvupadmin

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tvupadmin"
)

// YunosTvpubadminDeviceTvid 查询终端信息
// yunos.tvpubadmin.device.tvid
//
// 通过UUID查询终端信息
func YunosTvpubadminDeviceTvid(ctx context.Context, clt *core.SDKClient, req *tvupadmin.YunosTvpubadminDeviceTvidAPIRequest, resp *tvupadmin.YunosTvpubadminDeviceTvidAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
