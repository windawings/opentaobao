package alitripreceipt

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/alitripreceipt"
)

// AlitripReceiptSellerInvoiceReturn 飞猪发票商家回调接口
// alitrip.receipt.seller.invoice.return
//
// 飞猪发票回调接口
func AlitripReceiptSellerInvoiceReturn(ctx context.Context, clt *core.SDKClient, req *alitripreceipt.AlitripReceiptSellerInvoiceReturnAPIRequest, resp *alitripreceipt.AlitripReceiptSellerInvoiceReturnAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
