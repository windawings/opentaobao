package baichuan

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/baichuan"
)

// TaobaoBaichuanOpenaccountLogin 百川用户名密码登录
// taobao.baichuan.openaccount.login
//
// 百川用户名密码登录
func TaobaoBaichuanOpenaccountLogin(ctx context.Context, clt *core.SDKClient, req *baichuan.TaobaoBaichuanOpenaccountLoginAPIRequest, resp *baichuan.TaobaoBaichuanOpenaccountLoginAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
