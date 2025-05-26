package hotelhstdf

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/hotelhstdf"
)

// AlitripHotelHstdfPoilocationGet 根据平台城市id分页查询poi location
// alitrip.hotel.hstdf.poilocation.get
//
// 根据平台城市id分页查询poi location
func AlitripHotelHstdfPoilocationGet(ctx context.Context, clt *core.SDKClient, req *hotelhstdf.AlitripHotelHstdfPoilocationGetAPIRequest, resp *hotelhstdf.AlitripHotelHstdfPoilocationGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
