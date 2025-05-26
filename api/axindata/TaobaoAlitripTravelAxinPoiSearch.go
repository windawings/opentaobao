package axindata

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/axindata"
)

// TaobaoAlitripTravelAxinPoiSearch 景点poi搜索-阿信
// taobao.alitrip.travel.axin.poi.search
//
// 给阿信提供景点poi搜索
func TaobaoAlitripTravelAxinPoiSearch(ctx context.Context, clt *core.SDKClient, req *axindata.TaobaoAlitripTravelAxinPoiSearchAPIRequest, resp *axindata.TaobaoAlitripTravelAxinPoiSearchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
