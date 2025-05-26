package flight

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/flight"
)

// AlitripAgentFlightSellModifyList 销售改签单列表
// alitrip.agent.flight.sell.modify.list
//
// 销售改签单列表
func AlitripAgentFlightSellModifyList(ctx context.Context, clt *core.SDKClient, req *flight.AlitripAgentFlightSellModifyListAPIRequest, resp *flight.AlitripAgentFlightSellModifyListAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
