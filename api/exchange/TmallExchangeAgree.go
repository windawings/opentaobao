package exchange

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/exchange"
)

// TmallExchangeAgree 卖家同意换货申请
// tmall.exchange.agree
//
// 卖家同意换货申请
func TmallExchangeAgree(ctx context.Context, clt *core.SDKClient, req *exchange.TmallExchangeAgreeAPIRequest, resp *exchange.TmallExchangeAgreeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
