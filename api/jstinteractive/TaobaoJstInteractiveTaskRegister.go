package jstinteractive

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/jstinteractive"
)

// TaobaoJstInteractiveTaskRegister 互动任务开通接口
// taobao.jst.interactive.task.register
//
// 调用互动任务开通接口为小程序开通互动任务
func TaobaoJstInteractiveTaskRegister(ctx context.Context, clt *core.SDKClient, req *jstinteractive.TaobaoJstInteractiveTaskRegisterAPIRequest, resp *jstinteractive.TaobaoJstInteractiveTaskRegisterAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
