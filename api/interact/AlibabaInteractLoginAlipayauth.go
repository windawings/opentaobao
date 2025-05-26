package interact

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/interact"
)

// AlibabaInteractLoginAlipayauth 双11到店互动花呗红包获取token鉴权接口
// alibaba.interact.login.alipayauth
//
// 双11到店互动花呗红包获取token鉴权接口
func AlibabaInteractLoginAlipayauth(ctx context.Context, clt *core.SDKClient, req *interact.AlibabaInteractLoginAlipayauthAPIRequest, resp *interact.AlibabaInteractLoginAlipayauthAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
