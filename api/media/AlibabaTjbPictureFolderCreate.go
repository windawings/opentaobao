package media

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/media"
)

// AlibabaTjbPictureFolderCreate 淘特图片空间创建文件夹
// alibaba.tjb.picture.folder.create
//
// 淘特图片空间创建文件夹
func AlibabaTjbPictureFolderCreate(ctx context.Context, clt *core.SDKClient, req *media.AlibabaTjbPictureFolderCreateAPIRequest, resp *media.AlibabaTjbPictureFolderCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
