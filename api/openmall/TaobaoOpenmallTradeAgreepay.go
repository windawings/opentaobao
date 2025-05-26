package openmall

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/openmall"
)

// TaobaoOpenmallTradeAgreepay openmall订单支付
// taobao.openmall.trade.agreepay
//
// openmall订单支付
func TaobaoOpenmallTradeAgreepay(ctx context.Context, clt *core.SDKClient, req *openmall.TaobaoOpenmallTradeAgreepayAPIRequest, resp *openmall.TaobaoOpenmallTradeAgreepayAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
