package promotion

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/promotion"
)

// TaobaoPromotionCouponApply 优惠券领取
// taobao.promotion.coupon.apply
//
// 优惠券领取
func TaobaoPromotionCouponApply(ctx context.Context, clt *core.SDKClient, req *promotion.TaobaoPromotionCouponApplyAPIRequest, resp *promotion.TaobaoPromotionCouponApplyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
