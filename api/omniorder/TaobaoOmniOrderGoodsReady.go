package omniorder

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/omniorder"
)

// TaobaoOmniOrderGoodsReady 备货完成
// taobao.omni.order.goods.ready
//
// 备货完成
func TaobaoOmniOrderGoodsReady(ctx context.Context, clt *core.SDKClient, req *omniorder.TaobaoOmniOrderGoodsReadyAPIRequest, resp *omniorder.TaobaoOmniOrderGoodsReadyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
