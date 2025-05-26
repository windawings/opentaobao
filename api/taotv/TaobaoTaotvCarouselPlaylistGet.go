package taotv

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/taotv"
)

// TaobaoTaotvCarouselPlaylistGet 根据频道ID获取频道下节目单以及当前播放
// taobao.taotv.carousel.playlist.get
//
// 根据频道ID获取频道下节目单以及当前播放，包括所有视频源的视频
func TaobaoTaotvCarouselPlaylistGet(ctx context.Context, clt *core.SDKClient, req *taotv.TaobaoTaotvCarouselPlaylistGetAPIRequest, resp *taotv.TaobaoTaotvCarouselPlaylistGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
