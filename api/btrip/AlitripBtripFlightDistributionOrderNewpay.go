package btrip

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/btrip"
)

// AlitripBtripFlightDistributionOrderNewpay 商旅机票分销-订单支付V2
// alitrip.btrip.flight.distribution.order.newpay
//
// 商旅机票分销-订单支付V2
func AlitripBtripFlightDistributionOrderNewpay(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripFlightDistributionOrderNewpayAPIRequest, resp *btrip.AlitripBtripFlightDistributionOrderNewpayAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
