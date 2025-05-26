package ascp

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangWmsOrderCancel 回传发货单取消通知
// alibaba.dchain.aoxiang.wms.order.cancel
//
// 回传发货单取消通知
func AlibabaDchainAoxiangWmsOrderCancel(ctx context.Context, clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangWmsOrderCancelAPIRequest, resp *ascp.AlibabaDchainAoxiangWmsOrderCancelAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
