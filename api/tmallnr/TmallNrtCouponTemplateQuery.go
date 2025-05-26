package tmallnr

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tmallnr"
)

// TmallNrtCouponTemplateQuery 查询券模版
// tmall.nrt.coupon.template.query
//
// 查询券模版
func TmallNrtCouponTemplateQuery(ctx context.Context, clt *core.SDKClient, req *tmallnr.TmallNrtCouponTemplateQueryAPIRequest, resp *tmallnr.TmallNrtCouponTemplateQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
