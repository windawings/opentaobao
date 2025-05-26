package drugtrace

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugUpdatebarcodeBytraccode 根据追溯码修改69码
// alibaba.alihealth.drug.updatebarcode.bytraccode
//
// 根据追溯码修改69码
func AlibabaAlihealthDrugUpdatebarcodeBytraccode(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugUpdatebarcodeBytraccodeAPIRequest, resp *drugtrace.AlibabaAlihealthDrugUpdatebarcodeBytraccodeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
