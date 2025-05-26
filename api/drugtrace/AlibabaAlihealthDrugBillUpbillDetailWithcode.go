package drugtrace

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugBillUpbillDetailWithcode 查询上游出库单明细(带追溯码信息)
// alibaba.alihealth.drug.bill.upbill.detail.withcode
//
// 查询上游出库单明细(带追溯码信息)
func AlibabaAlihealthDrugBillUpbillDetailWithcode(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugBillUpbillDetailWithcodeAPIRequest, resp *drugtrace.AlibabaAlihealthDrugBillUpbillDetailWithcodeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
