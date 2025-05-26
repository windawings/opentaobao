package xhotelonlineorder

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/xhotelonlineorder"
)

// TaobaoXhotelFastinvoiceRequest 极速开票开票请求回传
// taobao.xhotel.fastinvoice.request
//
// 极速开票开票请求回传,用于记录航信开票请求数据
func TaobaoXhotelFastinvoiceRequest(ctx context.Context, clt *core.SDKClient, req *xhotelonlineorder.TaobaoXhotelFastinvoiceRequestAPIRequest, resp *xhotelonlineorder.TaobaoXhotelFastinvoiceRequestAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
