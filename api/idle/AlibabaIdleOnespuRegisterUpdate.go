package idle

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/idle"
)

// AlibabaIdleOnespuRegisterUpdate 闲鱼 ONESPU 挂载接口
// alibaba.idle.onespu.register.update
//
// 闲鱼 ONESPU 挂载接口
func AlibabaIdleOnespuRegisterUpdate(ctx context.Context, clt *core.SDKClient, req *idle.AlibabaIdleOnespuRegisterUpdateAPIRequest, resp *idle.AlibabaIdleOnespuRegisterUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
