package simba

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/simba"
)

// TaobaoSimbaCampaignBudgetGet 取得一个推广计划的日限额
// taobao.simba.campaign.budget.get
//
// 取得一个推广计划的日限额
func TaobaoSimbaCampaignBudgetGet(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaCampaignBudgetGetAPIRequest, resp *simba.TaobaoSimbaCampaignBudgetGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
