package examination

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/examination"
)

// AlibabaAlihealthExaminationReportDiagnoseOrderSummary 体检报告人工解读总结回传
// alibaba.alihealth.examination.report.diagnose.order.summary
//
// 记录体检报告人工解读总结
func AlibabaAlihealthExaminationReportDiagnoseOrderSummary(ctx context.Context, clt *core.SDKClient, req *examination.AlibabaAlihealthExaminationReportDiagnoseOrderSummaryAPIRequest, resp *examination.AlibabaAlihealthExaminationReportDiagnoseOrderSummaryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
