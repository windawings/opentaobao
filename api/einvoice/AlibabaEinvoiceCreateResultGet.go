package einvoice

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/einvoice"
)

// AlibabaEinvoiceCreateResultGet ERP开票结果获取
// alibaba.einvoice.create.result.get
//
// ERP开票结果获取
func AlibabaEinvoiceCreateResultGet(ctx context.Context, clt *core.SDKClient, req *einvoice.AlibabaEinvoiceCreateResultGetAPIRequest, resp *einvoice.AlibabaEinvoiceCreateResultGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
