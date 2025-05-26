package icbudropshipping

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/icbudropshipping"
)

// AlibabaDropshippingOrderPay alibaba dropshipping 支付代扣
// alibaba.dropshipping.order.pay
//
// alibaba dropshipping 支付代扣
func AlibabaDropshippingOrderPay(ctx context.Context, clt *core.SDKClient, req *icbudropshipping.AlibabaDropshippingOrderPayAPIRequest, resp *icbudropshipping.AlibabaDropshippingOrderPayAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
