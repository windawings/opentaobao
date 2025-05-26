package fenxiao

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/fenxiao"
)

// TaobaoFenxiaoDistributorItemsGet 查询商品下载记录
// taobao.fenxiao.distributor.items.get
//
// 供应商查询分销商商品下载记录。
func TaobaoFenxiaoDistributorItemsGet(ctx context.Context, clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoDistributorItemsGetAPIRequest, resp *fenxiao.TaobaoFenxiaoDistributorItemsGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
