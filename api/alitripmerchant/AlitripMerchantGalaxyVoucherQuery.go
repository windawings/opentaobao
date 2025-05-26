package alitripmerchant

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyVoucherQuery 查询单个代金券信息
// alitrip.merchant.galaxy.voucher.query
//
// 查询单个代金券信息
func AlitripMerchantGalaxyVoucherQuery(ctx context.Context, clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyVoucherQueryAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyVoucherQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
