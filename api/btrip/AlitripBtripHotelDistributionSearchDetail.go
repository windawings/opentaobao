package btrip

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/btrip"
)

// AlitripBtripHotelDistributionSearchDetail 商旅酒店api分销-详情报价接口
// alitrip.btrip.hotel.distribution.search.detail
//
// 商旅酒店api分销-详情报价接口
func AlitripBtripHotelDistributionSearchDetail(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripHotelDistributionSearchDetailAPIRequest, resp *btrip.AlitripBtripHotelDistributionSearchDetailAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
