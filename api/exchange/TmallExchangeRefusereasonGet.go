package exchange

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/exchange"
)

// TmallExchangeRefusereasonGet 获取拒绝换货原因列表
// tmall.exchange.refusereason.get
//
// 获取拒绝换货原因列表
func TmallExchangeRefusereasonGet(ctx context.Context, clt *core.SDKClient, req *exchange.TmallExchangeRefusereasonGetAPIRequest, resp *exchange.TmallExchangeRefusereasonGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
