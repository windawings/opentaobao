package qimen

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/qimen"
)

// TaobaoQimenShopSynchronize 店铺同步接口
// taobao.qimen.shop.synchronize
//
// 店铺同步接口描述
func TaobaoQimenShopSynchronize(ctx context.Context, clt *core.SDKClient, req *qimen.TaobaoQimenShopSynchronizeAPIRequest, resp *qimen.TaobaoQimenShopSynchronizeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
