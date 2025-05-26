package iot

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/iot"
)

// TaobaoAilabAicloudTopLikeList 列出收藏列表
// taobao.ailab.aicloud.top.like.list
//
// 列出收藏列表
func TaobaoAilabAicloudTopLikeList(ctx context.Context, clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopLikeListAPIRequest, resp *iot.TaobaoAilabAicloudTopLikeListAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
