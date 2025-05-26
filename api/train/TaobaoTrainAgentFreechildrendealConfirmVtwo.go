package train

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/train"
)

// TaobaoTrainAgentFreechildrendealConfirmVtwo 免费儿童处理
// taobao.train.agent.freechildrendeal.confirm.vtwo
//
// 免费儿童列表查询
func TaobaoTrainAgentFreechildrendealConfirmVtwo(ctx context.Context, clt *core.SDKClient, req *train.TaobaoTrainAgentFreechildrendealConfirmVtwoAPIRequest, resp *train.TaobaoTrainAgentFreechildrendealConfirmVtwoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
