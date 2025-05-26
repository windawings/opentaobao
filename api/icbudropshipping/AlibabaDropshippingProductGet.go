package icbudropshipping

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/icbudropshipping"
)

// AlibabaDropshippingProductGet 阿里巴巴dropshipping 产品信息获取
// alibaba.dropshipping.product.get
//
// 阿里巴巴dropshipping 产品信息获取
func AlibabaDropshippingProductGet(ctx context.Context, clt *core.SDKClient, req *icbudropshipping.AlibabaDropshippingProductGetAPIRequest, resp *icbudropshipping.AlibabaDropshippingProductGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
