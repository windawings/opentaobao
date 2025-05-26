package qimen

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/qimen"
)

// TaobaoQimenStockoutCreate 出库单创建接口
// taobao.qimen.stockout.create
//
// taobao.qimen.returnorder.create
func TaobaoQimenStockoutCreate(ctx context.Context, clt *core.SDKClient, req *qimen.TaobaoQimenStockoutCreateAPIRequest, resp *qimen.TaobaoQimenStockoutCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
