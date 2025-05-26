package tmallchannel

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tmallchannel"
)

// TmallChannelTradeOrderGets 分页查询采购单
// tmall.channel.trade.order.gets
//
// 分页查询采购单
func TmallChannelTradeOrderGets(ctx context.Context, clt *core.SDKClient, req *tmallchannel.TmallChannelTradeOrderGetsAPIRequest, resp *tmallchannel.TmallChannelTradeOrderGetsAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
