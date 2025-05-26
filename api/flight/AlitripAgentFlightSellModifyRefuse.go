package flight

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/flight"
)

// AlitripAgentFlightSellModifyRefuse 销售改签拒绝
// alitrip.agent.flight.sell.modify.refuse
//
// 销售改签拒绝
func AlitripAgentFlightSellModifyRefuse(ctx context.Context, clt *core.SDKClient, req *flight.AlitripAgentFlightSellModifyRefuseAPIRequest, resp *flight.AlitripAgentFlightSellModifyRefuseAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
