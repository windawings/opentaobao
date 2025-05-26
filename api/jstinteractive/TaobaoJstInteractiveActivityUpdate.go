package jstinteractive

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/jstinteractive"
)

// TaobaoJstInteractiveActivityUpdate 互动任务活动修改接口
// taobao.jst.interactive.activity.update
//
// 互动任务活动修改接口
func TaobaoJstInteractiveActivityUpdate(ctx context.Context, clt *core.SDKClient, req *jstinteractive.TaobaoJstInteractiveActivityUpdateAPIRequest, resp *jstinteractive.TaobaoJstInteractiveActivityUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
