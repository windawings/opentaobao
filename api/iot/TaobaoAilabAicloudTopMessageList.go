package iot

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/iot"
)

// TaobaoAilabAicloudTopMessageList 获取留言列表
// taobao.ailab.aicloud.top.message.list
//
// 根据指定参数获取留言列表
func TaobaoAilabAicloudTopMessageList(ctx context.Context, clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopMessageListAPIRequest, resp *iot.TaobaoAilabAicloudTopMessageListAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
