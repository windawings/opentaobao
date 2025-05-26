package baichuan

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/baichuan"
)

// TaobaoBaichuanOpenaccountRegister 百川账号注册
// taobao.baichuan.openaccount.register
//
// 百川账号注册
func TaobaoBaichuanOpenaccountRegister(ctx context.Context, clt *core.SDKClient, req *baichuan.TaobaoBaichuanOpenaccountRegisterAPIRequest, resp *baichuan.TaobaoBaichuanOpenaccountRegisterAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
