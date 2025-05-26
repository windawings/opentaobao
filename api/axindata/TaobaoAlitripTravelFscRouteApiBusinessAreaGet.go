package axindata

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/axindata"
)

// TaobaoAlitripTravelFscRouteApiBusinessAreaGet 获取业务区域
// taobao.alitrip.travel.fsc.route.api.business.area.get
//
// 获取业务区域
func TaobaoAlitripTravelFscRouteApiBusinessAreaGet(ctx context.Context, clt *core.SDKClient, req *axindata.TaobaoAlitripTravelFscRouteApiBusinessAreaGetAPIRequest, resp *axindata.TaobaoAlitripTravelFscRouteApiBusinessAreaGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
