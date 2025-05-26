package product

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/product"
)

// AlibabaJymItemExternalGoodsBatchDelete 交易猫外部商家批量删除商品接口
// alibaba.jym.item.external.goods.batch.delete
//
// 交易猫外部商家批量删除商品接口
func AlibabaJymItemExternalGoodsBatchDelete(ctx context.Context, clt *core.SDKClient, req *product.AlibabaJymItemExternalGoodsBatchDeleteAPIRequest, resp *product.AlibabaJymItemExternalGoodsBatchDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
