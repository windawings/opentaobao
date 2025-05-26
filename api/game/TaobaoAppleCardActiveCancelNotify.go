package game

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/game"
)

// TaobaoAppleCardActiveCancelNotify 苹果卡密取消激活回调接口
// taobao.apple.card.active.cancel.notify
//
// 苹果卡密取消激活回调接口
func TaobaoAppleCardActiveCancelNotify(ctx context.Context, clt *core.SDKClient, req *game.TaobaoAppleCardActiveCancelNotifyAPIRequest, resp *game.TaobaoAppleCardActiveCancelNotifyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
