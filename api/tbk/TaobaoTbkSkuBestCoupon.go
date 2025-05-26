package tbk

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tbk"
)

// TaobaoTbkSkuBestCoupon sku维度最优优惠券信息
// taobao.tbk.sku.best.coupon
//
// 根据itemid和skuid查询最优优惠券信息
func TaobaoTbkSkuBestCoupon(ctx context.Context, clt *core.SDKClient, req *tbk.TaobaoTbkSkuBestCouponAPIRequest, resp *tbk.TaobaoTbkSkuBestCouponAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
