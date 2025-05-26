package tbk

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tbk"
)

// TaobaoTbkShopGet 淘宝客-推广者-店铺搜索
// taobao.tbk.shop.get
//
// 淘宝客店铺查询
func TaobaoTbkShopGet(ctx context.Context, clt *core.SDKClient, req *tbk.TaobaoTbkShopGetAPIRequest, resp *tbk.TaobaoTbkShopGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
