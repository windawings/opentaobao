package simba

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/simba"
)

// TaobaoSimbaCampaignAreaoptionsGet 取得推广计划的可设置投放地域列表
// taobao.simba.campaign.areaoptions.get
//
// 取得推广计划的可设置投放地域列表
func TaobaoSimbaCampaignAreaoptionsGet(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaCampaignAreaoptionsGetAPIRequest, resp *simba.TaobaoSimbaCampaignAreaoptionsGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
