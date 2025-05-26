package tmallnr

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tmallnr"
)

// TmallNrtCouponInstanceQuery 查询用户券实例
// tmall.nrt.coupon.instance.query
//
// 查询用户券实例
func TmallNrtCouponInstanceQuery(ctx context.Context, clt *core.SDKClient, req *tmallnr.TmallNrtCouponInstanceQueryAPIRequest, resp *tmallnr.TmallNrtCouponInstanceQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
