package product

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/product"
)

// TmallItemSkuStatusGet 商品sku上下架查询
// tmall.item.sku.status.get
//
// 商品sku上下架状态查询
func TmallItemSkuStatusGet(ctx context.Context, clt *core.SDKClient, req *product.TmallItemSkuStatusGetAPIRequest, resp *product.TmallItemSkuStatusGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
