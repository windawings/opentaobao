package flightuppc

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/flightuppc"
)

// AlitripFlightInsuranceOrderSearch 查询保险订单详情
// alitrip.flight.insurance.order.search
//
// 查询保险订单详情
func AlitripFlightInsuranceOrderSearch(ctx context.Context, clt *core.SDKClient, req *flightuppc.AlitripFlightInsuranceOrderSearchAPIRequest, resp *flightuppc.AlitripFlightInsuranceOrderSearchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
