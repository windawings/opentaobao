package interact

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/interact"
)

// AlibabaInteractShopFavor 店铺收藏
// alibaba.interact.shop.favor
//
// 店铺收藏mtop接口开放鉴权接口，无入参出参，无安全风险，mtop接口开发 酒仙。
func AlibabaInteractShopFavor(ctx context.Context, clt *core.SDKClient, req *interact.AlibabaInteractShopFavorAPIRequest, resp *interact.AlibabaInteractShopFavorAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
