package opentrade

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/opentrade"
)

// TaobaoOpentradeToolsItemsBind 交易开放商品绑定
// taobao.opentrade.tools.items.bind
//
// 交易开放商品绑定
func TaobaoOpentradeToolsItemsBind(ctx context.Context, clt *core.SDKClient, req *opentrade.TaobaoOpentradeToolsItemsBindAPIRequest, resp *opentrade.TaobaoOpentradeToolsItemsBindAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
