package pur

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/pur"
)

// AlibabaPurSupplierInvoicecreate preInvoice创建
// alibaba.pur.supplier.invoicecreate
//
// preInvoice创建
func AlibabaPurSupplierInvoicecreate(ctx context.Context, clt *core.SDKClient, req *pur.AlibabaPurSupplierInvoicecreateAPIRequest, resp *pur.AlibabaPurSupplierInvoicecreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
