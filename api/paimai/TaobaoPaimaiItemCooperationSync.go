package paimai

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/paimai"
)

// TaobaoPaimaiItemCooperationSync 商品同步
// taobao.paimai.item.cooperation.sync
//
// 商品同步
func TaobaoPaimaiItemCooperationSync(ctx context.Context, clt *core.SDKClient, req *paimai.TaobaoPaimaiItemCooperationSyncAPIRequest, resp *paimai.TaobaoPaimaiItemCooperationSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
