package alitripmerchant

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyOrderQueryInfo 订单详情改版
// alitrip.merchant.galaxy.order.query.info
//
// 订单页详情查询
func AlitripMerchantGalaxyOrderQueryInfo(ctx context.Context, clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyOrderQueryInfoAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyOrderQueryInfoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
