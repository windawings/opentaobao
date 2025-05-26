package flight

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/flight"
)

// AlitripAgentFlightSellRefundDetail 销售退票单详情
// alitrip.agent.flight.sell.refund.detail
//
// 销售退票单详情
func AlitripAgentFlightSellRefundDetail(ctx context.Context, clt *core.SDKClient, req *flight.AlitripAgentFlightSellRefundDetailAPIRequest, resp *flight.AlitripAgentFlightSellRefundDetailAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
