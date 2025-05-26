package xhotel

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/xhotel"
)

// TaobaoXhotelDataServiceHotelServiceindex 酒店服务指数
// taobao.xhotel.data.service.hotel.serviceindex
//
// 酒店服务指数
func TaobaoXhotelDataServiceHotelServiceindex(ctx context.Context, clt *core.SDKClient, req *xhotel.TaobaoXhotelDataServiceHotelServiceindexAPIRequest, resp *xhotel.TaobaoXhotelDataServiceHotelServiceindexAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
