package axindata

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/axindata"
)

// TaobaoAlitripTravelFscRouteApiPoiGet 获取景点（POI）信息
// taobao.alitrip.travel.fsc.route.api.poi.get
//
// 获取景点（POI）信息
func TaobaoAlitripTravelFscRouteApiPoiGet(ctx context.Context, clt *core.SDKClient, req *axindata.TaobaoAlitripTravelFscRouteApiPoiGetAPIRequest, resp *axindata.TaobaoAlitripTravelFscRouteApiPoiGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
