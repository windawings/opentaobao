package flight

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/flight"
)

// AlitripPolicyDomfareFlowdata 比价工具流量详情
// alitrip.policy.domfare.flowdata
//
// 比价工具流量详情
func AlitripPolicyDomfareFlowdata(ctx context.Context, clt *core.SDKClient, req *flight.AlitripPolicyDomfareFlowdataAPIRequest, resp *flight.AlitripPolicyDomfareFlowdataAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
