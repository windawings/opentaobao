package drugtrace

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytScqyUploadrelation 关联关系上传
// alibaba.alihealth.drug.kyt.scqy.uploadrelation
//
// 关联关系上传
func AlibabaAlihealthDrugKytScqyUploadrelation(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytScqyUploadrelationAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytScqyUploadrelationAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
