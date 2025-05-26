package exchange

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/exchange"
)

// TmallExchangeRefuse 卖家拒绝换货申请
// tmall.exchange.refuse
//
// 卖家拒绝换货申请
func TmallExchangeRefuse(ctx context.Context, clt *core.SDKClient, req *exchange.TmallExchangeRefuseAPIRequest, resp *exchange.TmallExchangeRefuseAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
