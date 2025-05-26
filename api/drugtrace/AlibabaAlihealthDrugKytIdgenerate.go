package drugtrace

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytIdgenerate 终端(医疗机构|零售药店)ID生成接口
// alibaba.alihealth.drug.kyt.idgenerate
//
// 终端(医疗机构|零售药店)ID生成接口
func AlibabaAlihealthDrugKytIdgenerate(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytIdgenerateAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytIdgenerateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
