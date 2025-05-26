package einvoice

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/einvoice"
)

// AlibabaEinvoiceOrderRefundUpdate 回传订单退款审核结果
// alibaba.einvoice.order.refund.update
//
// ISV回传订单退款审核结果
func AlibabaEinvoiceOrderRefundUpdate(ctx context.Context, clt *core.SDKClient, req *einvoice.AlibabaEinvoiceOrderRefundUpdateAPIRequest, resp *einvoice.AlibabaEinvoiceOrderRefundUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
