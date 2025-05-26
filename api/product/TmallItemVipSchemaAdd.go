package product

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/product"
)

// TmallItemVipSchemaAdd 大商家商品发布接口
// tmall.item.vip.schema.add
//
// 大商家商品发布接口
func TmallItemVipSchemaAdd(ctx context.Context, clt *core.SDKClient, req *product.TmallItemVipSchemaAddAPIRequest, resp *product.TmallItemVipSchemaAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
