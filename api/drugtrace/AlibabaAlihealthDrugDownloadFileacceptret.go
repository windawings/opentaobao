package drugtrace

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugDownloadFileacceptret 企业上传回执
// alibaba.alihealth.drug.download.fileacceptret
//
// 拿到企业下载回执，将企业已下载的和未下载成功的条目都相应的改变状态
func AlibabaAlihealthDrugDownloadFileacceptret(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugDownloadFileacceptretAPIRequest, resp *drugtrace.AlibabaAlihealthDrugDownloadFileacceptretAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
