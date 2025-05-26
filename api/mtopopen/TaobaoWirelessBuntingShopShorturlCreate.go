package mtopopen

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/mtopopen"
)

// TaobaoWirelessBuntingShopShorturlCreate 通过店铺id取得短链
// taobao.wireless.bunting.shop.shorturl.create
//
// 通过店铺id取得短链
func TaobaoWirelessBuntingShopShorturlCreate(ctx context.Context, clt *core.SDKClient, req *mtopopen.TaobaoWirelessBuntingShopShorturlCreateAPIRequest, resp *mtopopen.TaobaoWirelessBuntingShopShorturlCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
