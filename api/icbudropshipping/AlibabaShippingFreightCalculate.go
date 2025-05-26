package icbudropshipping

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/icbudropshipping"
)

// AlibabaShippingFreightCalculate 阿里巴巴商品运费计算查询接口
// alibaba.shipping.freight.calculate
//
// 阿里巴巴商品运费计算查询接口
func AlibabaShippingFreightCalculate(ctx context.Context, clt *core.SDKClient, req *icbudropshipping.AlibabaShippingFreightCalculateAPIRequest, resp *icbudropshipping.AlibabaShippingFreightCalculateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
