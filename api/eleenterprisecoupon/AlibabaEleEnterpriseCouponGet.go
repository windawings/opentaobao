package eleenterprisecoupon

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/eleenterprisecoupon"
)

// AlibabaEleEnterpriseCouponGet 获取用户优惠券
// alibaba.ele.enterprise.coupon.get
//
// 获取用户优惠券
func AlibabaEleEnterpriseCouponGet(ctx context.Context, clt *core.SDKClient, req *eleenterprisecoupon.AlibabaEleEnterpriseCouponGetAPIRequest, resp *eleenterprisecoupon.AlibabaEleEnterpriseCouponGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
