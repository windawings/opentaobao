package drugtrace

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytUpdatebillinfo 零售端平台单据更新
// alibaba.alihealth.drug.kyt.updatebillinfo
//
// 零售端平台单据更新
func AlibabaAlihealthDrugKytUpdatebillinfo(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytUpdatebillinfoAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytUpdatebillinfoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
