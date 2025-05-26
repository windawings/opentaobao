package mos

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/mos"
)

// AlibabaMosHrBackgroundReportNotify 背调公司背调结果通知
// alibaba.mos.hr.background.report.notify
//
// 背调公司背调结果通知
func AlibabaMosHrBackgroundReportNotify(ctx context.Context, clt *core.SDKClient, req *mos.AlibabaMosHrBackgroundReportNotifyAPIRequest, resp *mos.AlibabaMosHrBackgroundReportNotifyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
