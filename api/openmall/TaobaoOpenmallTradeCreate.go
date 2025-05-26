package openmall

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/openmall"
)

// TaobaoOpenmallTradeCreate 创建订单
// taobao.openmall.trade.create
//
// 创建Openmall订单
func TaobaoOpenmallTradeCreate(ctx context.Context, clt *core.SDKClient, req *openmall.TaobaoOpenmallTradeCreateAPIRequest, resp *openmall.TaobaoOpenmallTradeCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
