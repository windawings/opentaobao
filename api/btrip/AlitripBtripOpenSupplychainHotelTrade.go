package btrip

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/btrip"
)

// AlitripBtripOpenSupplychainHotelTrade 【商旅】酒店交易查询流水接口
// alitrip.btrip.open.supplychain.hotel.trade
//
// 【商旅】酒店交易查询流水接口——杭州市政府
func AlitripBtripOpenSupplychainHotelTrade(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripOpenSupplychainHotelTradeAPIRequest, resp *btrip.AlitripBtripOpenSupplychainHotelTradeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
