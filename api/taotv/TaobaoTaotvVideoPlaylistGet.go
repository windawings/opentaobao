package taotv

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/taotv"
)

// TaobaoTaotvVideoPlaylistGet 根据频道ID获取频道下节目单以及当前播放
// taobao.taotv.video.playlist.get
//
// 根据频道ID获取频道下节目单以及当前播放
func TaobaoTaotvVideoPlaylistGet(ctx context.Context, clt *core.SDKClient, req *taotv.TaobaoTaotvVideoPlaylistGetAPIRequest, resp *taotv.TaobaoTaotvVideoPlaylistGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
