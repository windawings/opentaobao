package wms

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/wms"
)

// TaobaoWlbWmsOrderCancelNotify 单据取消接口
// taobao.wlb.wms.order.cancel.notify
//
// 单据取消接口
func TaobaoWlbWmsOrderCancelNotify(ctx context.Context, clt *core.SDKClient, req *wms.TaobaoWlbWmsOrderCancelNotifyAPIRequest, resp *wms.TaobaoWlbWmsOrderCancelNotifyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
