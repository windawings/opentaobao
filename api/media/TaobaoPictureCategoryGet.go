package media

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/media"
)

// TaobaoPictureCategoryGet 获取图片分类信息
// taobao.picture.category.get
//
// 获取图片分类信息
func TaobaoPictureCategoryGet(ctx context.Context, clt *core.SDKClient, req *media.TaobaoPictureCategoryGetAPIRequest, resp *media.TaobaoPictureCategoryGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
