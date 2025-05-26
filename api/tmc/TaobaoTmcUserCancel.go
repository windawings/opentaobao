package tmc

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tmc"
)

// TaobaoTmcUserCancel 取消用户的消息服务
// taobao.tmc.user.cancel
//
// 取消用户的消息服务
func TaobaoTmcUserCancel(ctx context.Context, clt *core.SDKClient, req *tmc.TaobaoTmcUserCancelAPIRequest, resp *tmc.TaobaoTmcUserCancelAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
