package fundplatform

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/fundplatform"
)

// AlibabaCfoIncomingInvoiceLedgerFullysync 票易通全量底账数据同步
// alibaba.cfo.incoming.invoice.ledger.fullysync
//
// 票易通全量底账数据同步
func AlibabaCfoIncomingInvoiceLedgerFullysync(ctx context.Context, clt *core.SDKClient, req *fundplatform.AlibabaCfoIncomingInvoiceLedgerFullysyncAPIRequest, resp *fundplatform.AlibabaCfoIncomingInvoiceLedgerFullysyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
