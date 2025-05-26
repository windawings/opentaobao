package fenxiao

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/fenxiao"
)

// TaobaoFenxiaoProductSkuUpdate 产品sku编辑接口
// taobao.fenxiao.product.sku.update
//
// 产品SKU信息更新
func TaobaoFenxiaoProductSkuUpdate(ctx context.Context, clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoProductSkuUpdateAPIRequest, resp *fenxiao.TaobaoFenxiaoProductSkuUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
