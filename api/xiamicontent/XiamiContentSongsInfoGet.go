package xiamicontent

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/xiamicontent"
)

// XiamiContentSongsInfoGet 获取歌曲信息
// xiami.content.songs.info.get
//
// (批量)获取歌曲信息
func XiamiContentSongsInfoGet(ctx context.Context, clt *core.SDKClient, req *xiamicontent.XiamiContentSongsInfoGetAPIRequest, resp *xiamicontent.XiamiContentSongsInfoGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
