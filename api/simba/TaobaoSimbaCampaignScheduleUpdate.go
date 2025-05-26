package simba

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/simba"
)

// TaobaoSimbaCampaignScheduleUpdate 更新一个推广计划的分时折扣设置
// taobao.simba.campaign.schedule.update
//
// 更新一个推广计划的分时折扣设置
func TaobaoSimbaCampaignScheduleUpdate(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaCampaignScheduleUpdateAPIRequest, resp *simba.TaobaoSimbaCampaignScheduleUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
