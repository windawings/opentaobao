package flight

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/flight"
)

// AlitripAgentCoordinateHandle 慧飞商家协同单接手接口
// alitrip.agent.coordinate.handle
//
// 慧飞商家协同单接手接口
func AlitripAgentCoordinateHandle(ctx context.Context, clt *core.SDKClient, req *flight.AlitripAgentCoordinateHandleAPIRequest, resp *flight.AlitripAgentCoordinateHandleAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
