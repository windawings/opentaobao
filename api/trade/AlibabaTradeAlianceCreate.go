package trade

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/trade"
)

// AlibabaTradeAlianceCreate 推客平台订单回流
// alibaba.trade.aliance.create
//
// 推客平台订单回流
func AlibabaTradeAlianceCreate(ctx context.Context, clt *core.SDKClient, req *trade.AlibabaTradeAlianceCreateAPIRequest, resp *trade.AlibabaTradeAlianceCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
