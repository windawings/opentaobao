package exchange

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/exchange"
)

// TmallExchangeGet 获取单笔换货详情
// tmall.exchange.get
//
// 获取单笔换货详情
func TmallExchangeGet(ctx context.Context, clt *core.SDKClient, req *exchange.TmallExchangeGetAPIRequest, resp *exchange.TmallExchangeGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
