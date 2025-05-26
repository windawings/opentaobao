package alitripmerchant

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyDerbyMemberVoucherCardApsRefund Aps退券通知接口
// alitrip.merchant.galaxy.derby.member.voucher.card.aps.refund
//
// Aps退券通知接口
func AlitripMerchantGalaxyDerbyMemberVoucherCardApsRefund(ctx context.Context, clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherCardApsRefundAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherCardApsRefundAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
