package drugtrace

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytCodetobill 通过追溯码查单据
// alibaba.alihealth.drug.kyt.codetobill
//
// 通过追溯码查单据
func AlibabaAlihealthDrugKytCodetobill(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytCodetobillAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytCodetobillAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
