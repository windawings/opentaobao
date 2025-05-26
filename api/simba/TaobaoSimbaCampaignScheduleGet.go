package simba

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/simba"
)

// TaobaoSimbaCampaignScheduleGet 取得一个推广计划的分时折扣设置
// taobao.simba.campaign.schedule.get
//
// 取得一个推广计划的分时折扣设置
func TaobaoSimbaCampaignScheduleGet(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaCampaignScheduleGetAPIRequest, resp *simba.TaobaoSimbaCampaignScheduleGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
