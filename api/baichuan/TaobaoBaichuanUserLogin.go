package baichuan

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/baichuan"
)

// TaobaoBaichuanUserLogin 百川H5登录
// taobao.baichuan.user.login
//
// 百川H5登录
func TaobaoBaichuanUserLogin(ctx context.Context, clt *core.SDKClient, req *baichuan.TaobaoBaichuanUserLoginAPIRequest, resp *baichuan.TaobaoBaichuanUserLoginAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
