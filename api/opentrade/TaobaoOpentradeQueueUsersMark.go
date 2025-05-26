package opentrade

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/opentrade"
)

// TaobaoOpentradeQueueUsersMark 尖货交易可购买用户标记
// taobao.opentrade.queue.users.mark
//
// 尖货交易用户标记信息回传，根据openId标记用户可购买商品
func TaobaoOpentradeQueueUsersMark(ctx context.Context, clt *core.SDKClient, req *opentrade.TaobaoOpentradeQueueUsersMarkAPIRequest, resp *opentrade.TaobaoOpentradeQueueUsersMarkAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
