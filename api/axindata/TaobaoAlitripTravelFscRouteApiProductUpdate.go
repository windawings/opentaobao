package axindata

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/axindata"
)

// TaobaoAlitripTravelFscRouteApiProductUpdate 更新线路产品基本信息
// taobao.alitrip.travel.fsc.route.api.product.update
//
// 更新线路产品基本信息
func TaobaoAlitripTravelFscRouteApiProductUpdate(ctx context.Context, clt *core.SDKClient, req *axindata.TaobaoAlitripTravelFscRouteApiProductUpdateAPIRequest, resp *axindata.TaobaoAlitripTravelFscRouteApiProductUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
