package tvupadmin

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tvupadmin"
)

// YunosTvpubadminContentTopicOffline 迎客松专题下线
// yunos.tvpubadmin.content.topic.offline
//
// 迎客松专题下线
func YunosTvpubadminContentTopicOffline(ctx context.Context, clt *core.SDKClient, req *tvupadmin.YunosTvpubadminContentTopicOfflineAPIRequest, resp *tvupadmin.YunosTvpubadminContentTopicOfflineAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
