package drugtrace

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugGetbarcodeBytraccode 根据追溯码获取69码
// alibaba.alihealth.drug.getbarcode.bytraccode
//
// 根据追溯码获取69码
func AlibabaAlihealthDrugGetbarcodeBytraccode(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugGetbarcodeBytraccodeAPIRequest, resp *drugtrace.AlibabaAlihealthDrugGetbarcodeBytraccodeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
