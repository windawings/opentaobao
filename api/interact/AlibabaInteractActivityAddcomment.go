package interact

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/interact"
)

// AlibabaInteractActivityAddcomment 微淘评论接口
// alibaba.interact.activity.addcomment
//
// 发表评论，并返回楼层
func AlibabaInteractActivityAddcomment(ctx context.Context, clt *core.SDKClient, req *interact.AlibabaInteractActivityAddcommentAPIRequest, resp *interact.AlibabaInteractActivityAddcommentAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
