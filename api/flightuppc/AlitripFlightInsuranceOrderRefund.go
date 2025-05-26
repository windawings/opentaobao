package flightuppc

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/flightuppc"
)

// AlitripFlightInsuranceOrderRefund 保险订单退保
// alitrip.flight.insurance.order.refund
//
// 保险订单退保
func AlitripFlightInsuranceOrderRefund(ctx context.Context, clt *core.SDKClient, req *flightuppc.AlitripFlightInsuranceOrderRefundAPIRequest, resp *flightuppc.AlitripFlightInsuranceOrderRefundAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
