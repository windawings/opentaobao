package alihouse

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/alihouse"
)

// TmallAlihouseTradeCouponOrderStatusQuery 查询电商券履约单状态
// tmall.alihouse.trade.coupon.order.status.query
//
// 查询电商券履约单状态
func TmallAlihouseTradeCouponOrderStatusQuery(ctx context.Context, clt *core.SDKClient, req *alihouse.TmallAlihouseTradeCouponOrderStatusQueryAPIRequest, resp *alihouse.TmallAlihouseTradeCouponOrderStatusQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
