package miniappopen

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/miniappopen"
)

// TaobaoMiniappVirtualItemGet 小程序关联虚拟商品查询
// taobao.miniapp.virtual.item.get
//
// 小程序关联虚拟商品查询
func TaobaoMiniappVirtualItemGet(ctx context.Context, clt *core.SDKClient, req *miniappopen.TaobaoMiniappVirtualItemGetAPIRequest, resp *miniappopen.TaobaoMiniappVirtualItemGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
