package baichuan

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/baichuan"
)

// TaobaoBaichuanOpenaccountResetcodeSend 百川发送找回密码验证码
// taobao.baichuan.openaccount.resetcode.send
//
// 百川发送找回密码验证码
func TaobaoBaichuanOpenaccountResetcodeSend(ctx context.Context, clt *core.SDKClient, req *baichuan.TaobaoBaichuanOpenaccountResetcodeSendAPIRequest, resp *baichuan.TaobaoBaichuanOpenaccountResetcodeSendAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
