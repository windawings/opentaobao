package btrip

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/btrip"
)

// AlitripBtripFlightDistributionRefundNewdetail 商旅机票退票详情接口V2
// alitrip.btrip.flight.distribution.refund.newdetail
//
// 商旅机票退票详情接口V2
func AlitripBtripFlightDistributionRefundNewdetail(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripFlightDistributionRefundNewdetailAPIRequest, resp *btrip.AlitripBtripFlightDistributionRefundNewdetailAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
