package media

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/media"
)

// TaobaoPictureGet 获取图片信息
// taobao.picture.get
//
// 获取图片信息
func TaobaoPictureGet(ctx context.Context, clt *core.SDKClient, req *media.TaobaoPictureGetAPIRequest, resp *media.TaobaoPictureGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
