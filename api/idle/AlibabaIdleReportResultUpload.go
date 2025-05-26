package idle

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/idle"
)

// AlibabaIdleReportResultUpload 服务商上传验货报告
// alibaba.idle.report.result.upload
//
// 服务商上传验货报告
func AlibabaIdleReportResultUpload(ctx context.Context, clt *core.SDKClient, req *idle.AlibabaIdleReportResultUploadAPIRequest, resp *idle.AlibabaIdleReportResultUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
