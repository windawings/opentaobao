package tvupadmin

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tvupadmin"
)

// YunosTvpubadminContentChannelOffline 迎客松影视频道下线
// yunos.tvpubadmin.content.channel.offline
//
// 迎客松影视频道下线
func YunosTvpubadminContentChannelOffline(ctx context.Context, clt *core.SDKClient, req *tvupadmin.YunosTvpubadminContentChannelOfflineAPIRequest, resp *tvupadmin.YunosTvpubadminContentChannelOfflineAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
