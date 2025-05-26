package media

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/media"
)

// TaobaoMediaVideoList 获取商家视频列表
// taobao.media.video.list
//
// 用于获取授权商家的视频列表
func TaobaoMediaVideoList(ctx context.Context, clt *core.SDKClient, req *media.TaobaoMediaVideoListAPIRequest, resp *media.TaobaoMediaVideoListAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
