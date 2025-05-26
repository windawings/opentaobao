package xhotel

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/xhotel"
)

// TaobaoXhotelOrderOfflineSettleCancel 线下信用住取消结账专用接口
// taobao.xhotel.order.offline.settle.cancel
//
// 线下信用住取消结账专用接口
func TaobaoXhotelOrderOfflineSettleCancel(ctx context.Context, clt *core.SDKClient, req *xhotel.TaobaoXhotelOrderOfflineSettleCancelAPIRequest, resp *xhotel.TaobaoXhotelOrderOfflineSettleCancelAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
