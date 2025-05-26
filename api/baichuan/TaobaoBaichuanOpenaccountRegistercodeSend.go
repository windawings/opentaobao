package baichuan

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/baichuan"
)

// TaobaoBaichuanOpenaccountRegistercodeSend 百川发送注册验证码
// taobao.baichuan.openaccount.registercode.send
//
// 百川发送注册验证码
func TaobaoBaichuanOpenaccountRegistercodeSend(ctx context.Context, clt *core.SDKClient, req *baichuan.TaobaoBaichuanOpenaccountRegistercodeSendAPIRequest, resp *baichuan.TaobaoBaichuanOpenaccountRegistercodeSendAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
