package wms

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/wms"
)

// TaobaoWlbWmsReturnOrderNotify 销售退货通知
// taobao.wlb.wms.return.order.notify
//
// 销售退货通知
func TaobaoWlbWmsReturnOrderNotify(ctx context.Context, clt *core.SDKClient, req *wms.TaobaoWlbWmsReturnOrderNotifyAPIRequest, resp *wms.TaobaoWlbWmsReturnOrderNotifyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
