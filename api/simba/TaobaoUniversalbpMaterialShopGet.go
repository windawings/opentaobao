package simba

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/simba"
)

// TaobaoUniversalbpMaterialShopGet 获取店铺信息
// taobao.universalbp.material.shop.get
//
// 获取店铺信息
func TaobaoUniversalbpMaterialShopGet(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoUniversalbpMaterialShopGetAPIRequest, resp *simba.TaobaoUniversalbpMaterialShopGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
