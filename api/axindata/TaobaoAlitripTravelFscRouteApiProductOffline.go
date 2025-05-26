package axindata

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/axindata"
)

// TaobaoAlitripTravelFscRouteApiProductOffline 产品下线
// taobao.alitrip.travel.fsc.route.api.product.offline
//
// 产品下线
func TaobaoAlitripTravelFscRouteApiProductOffline(ctx context.Context, clt *core.SDKClient, req *axindata.TaobaoAlitripTravelFscRouteApiProductOfflineAPIRequest, resp *axindata.TaobaoAlitripTravelFscRouteApiProductOfflineAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
