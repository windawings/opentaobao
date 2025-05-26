package flight

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/flight"
)

// AlitripAgentFlightSellModifyBackfill 销售改签回填
// alitrip.agent.flight.sell.modify.backfill
//
// 销售改签回填
func AlitripAgentFlightSellModifyBackfill(ctx context.Context, clt *core.SDKClient, req *flight.AlitripAgentFlightSellModifyBackfillAPIRequest, resp *flight.AlitripAgentFlightSellModifyBackfillAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
