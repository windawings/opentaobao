package omniorder

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/omniorder"
)

// TaobaoOmniitemSkuGet 获取全渠道门店商品sku
// taobao.omniitem.sku.get
//
// 通过skuId或者skuOutId查询全渠道门店商品sku信息
func TaobaoOmniitemSkuGet(ctx context.Context, clt *core.SDKClient, req *omniorder.TaobaoOmniitemSkuGetAPIRequest, resp *omniorder.TaobaoOmniitemSkuGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
