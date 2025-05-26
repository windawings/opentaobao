package txcs

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/txcs"
)

// TmallTxcsFinanceVerifyStatementBill 供应商核销单录入
// tmall.txcs.finance.verify.statement.bill
//
// 供应商核销单录入
func TmallTxcsFinanceVerifyStatementBill(ctx context.Context, clt *core.SDKClient, req *txcs.TmallTxcsFinanceVerifyStatementBillAPIRequest, resp *txcs.TmallTxcsFinanceVerifyStatementBillAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
