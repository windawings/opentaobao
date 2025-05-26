package game

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/game"
)

// TaobaoAppleNewuserSignNotifyNewversion 新用户签约结果通知接口v2
// taobao.apple.newuser.sign.notify.newversion
//
// 资和信主动通知签约结果
func TaobaoAppleNewuserSignNotifyNewversion(ctx context.Context, clt *core.SDKClient, req *game.TaobaoAppleNewuserSignNotifyNewversionAPIRequest, resp *game.TaobaoAppleNewuserSignNotifyNewversionAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
