package product

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/product"
)

// TmallItemVipUpdateSchemaGet vip商家编辑商品的规则获取接口
// tmall.item.vip.update.schema.get
//
// 获取vip商家编辑商品的规则
func TmallItemVipUpdateSchemaGet(ctx context.Context, clt *core.SDKClient, req *product.TmallItemVipUpdateSchemaGetAPIRequest, resp *product.TmallItemVipUpdateSchemaGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
