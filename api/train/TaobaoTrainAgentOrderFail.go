package train

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/train"
)

// TaobaoTrainAgentOrderFail 出票失败
// taobao.train.agent.order.fail
//
// 出票失败
func TaobaoTrainAgentOrderFail(ctx context.Context, clt *core.SDKClient, req *train.TaobaoTrainAgentOrderFailAPIRequest, resp *train.TaobaoTrainAgentOrderFailAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
