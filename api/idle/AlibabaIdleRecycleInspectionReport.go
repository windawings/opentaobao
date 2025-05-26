package idle

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/idle"
)

// AlibabaIdleRecycleInspectionReport 鉴定报告
// alibaba.idle.recycle.inspection.report
//
// 回收商鉴定报告
func AlibabaIdleRecycleInspectionReport(ctx context.Context, clt *core.SDKClient, req *idle.AlibabaIdleRecycleInspectionReportAPIRequest, resp *idle.AlibabaIdleRecycleInspectionReportAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
