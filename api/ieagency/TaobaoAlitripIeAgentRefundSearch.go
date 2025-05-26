package ieagency

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/ieagency"
)

// TaobaoAlitripIeAgentRefundSearch 卖家查询退票申请
// taobao.alitrip.ie.agent.refund.search
//
// 卖家查询退票申请
func TaobaoAlitripIeAgentRefundSearch(ctx context.Context, clt *core.SDKClient, req *ieagency.TaobaoAlitripIeAgentRefundSearchAPIRequest, resp *ieagency.TaobaoAlitripIeAgentRefundSearchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
