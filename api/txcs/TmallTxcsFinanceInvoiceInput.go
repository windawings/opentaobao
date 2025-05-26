package txcs

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/txcs"
)

// TmallTxcsFinanceInvoiceInput 供应商发票录入
// tmall.txcs.finance.invoice.input
//
// 提供天猫超市外部合作商家财务：供应商发票录入
func TmallTxcsFinanceInvoiceInput(ctx context.Context, clt *core.SDKClient, req *txcs.TmallTxcsFinanceInvoiceInputAPIRequest, resp *txcs.TmallTxcsFinanceInvoiceInputAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
