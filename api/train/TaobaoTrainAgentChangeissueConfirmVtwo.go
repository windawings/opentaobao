package train

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/train"
)

// TaobaoTrainAgentChangeissueConfirmVtwo 火车票代理商接口-跑腿改签出票回填-含鉴权校验
// taobao.train.agent.changeissue.confirm.vtwo
//
// 火车票代理商接口-跑腿改签出票回填-含鉴权校验
func TaobaoTrainAgentChangeissueConfirmVtwo(ctx context.Context, clt *core.SDKClient, req *train.TaobaoTrainAgentChangeissueConfirmVtwoAPIRequest, resp *train.TaobaoTrainAgentChangeissueConfirmVtwoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
