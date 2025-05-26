package paimai

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/paimai"
)

// TaobaoAuctionBeikeItemSync 贝壳商品同步接口
// taobao.auction.beike.item.sync
//
// 贝壳商品同步接口
func TaobaoAuctionBeikeItemSync(ctx context.Context, clt *core.SDKClient, req *paimai.TaobaoAuctionBeikeItemSyncAPIRequest, resp *paimai.TaobaoAuctionBeikeItemSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
