package fenxiao

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/fenxiao"
)

// TaobaoFenxiaoProductSkuDelete 产品SKU删除接口
// taobao.fenxiao.product.sku.delete
//
// 根据sku properties删除sku数据
func TaobaoFenxiaoProductSkuDelete(ctx context.Context, clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoProductSkuDeleteAPIRequest, resp *fenxiao.TaobaoFenxiaoProductSkuDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
