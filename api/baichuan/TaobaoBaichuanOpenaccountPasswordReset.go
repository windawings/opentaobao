package baichuan

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/baichuan"
)

// TaobaoBaichuanOpenaccountPasswordReset 百川找回密码
// taobao.baichuan.openaccount.password.reset
//
// 百川找回密码
func TaobaoBaichuanOpenaccountPasswordReset(ctx context.Context, clt *core.SDKClient, req *baichuan.TaobaoBaichuanOpenaccountPasswordResetAPIRequest, resp *baichuan.TaobaoBaichuanOpenaccountPasswordResetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
