package fenxiao

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/fenxiao"
)

// TaobaoFenxiaoProductSkusGet SKU查询接口
// taobao.fenxiao.product.skus.get
//
// 产品sku查询
func TaobaoFenxiaoProductSkusGet(ctx context.Context, clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoProductSkusGetAPIRequest, resp *fenxiao.TaobaoFenxiaoProductSkusGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
