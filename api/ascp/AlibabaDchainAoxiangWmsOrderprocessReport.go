package ascp

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangWmsOrderprocessReport 回传发货单流水通知
// alibaba.dchain.aoxiang.wms.orderprocess.report
//
// 回传发货单流水通知
func AlibabaDchainAoxiangWmsOrderprocessReport(ctx context.Context, clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangWmsOrderprocessReportAPIRequest, resp *ascp.AlibabaDchainAoxiangWmsOrderprocessReportAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
