package product

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/product"
)

// TmallItemStoreSchemaUpdate 天猫门店商品编辑
// tmall.item.store.schema.update
//
// 天猫门店商品编辑
func TmallItemStoreSchemaUpdate(ctx context.Context, clt *core.SDKClient, req *product.TmallItemStoreSchemaUpdateAPIRequest, resp *product.TmallItemStoreSchemaUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
