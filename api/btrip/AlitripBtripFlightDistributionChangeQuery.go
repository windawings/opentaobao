package btrip

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/btrip"
)

// AlitripBtripFlightDistributionChangeQuery 改签航班询价
// alitrip.btrip.flight.distribution.change.query
//
// 商旅机票改签航班询价
func AlitripBtripFlightDistributionChangeQuery(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripFlightDistributionChangeQueryAPIRequest, resp *btrip.AlitripBtripFlightDistributionChangeQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
