package simba

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/simba"
)

// TaobaoSubwayCampaignOfflineFind 查询某计划离线多日汇总列表
// taobao.subway.campaign.offline.find
//
// 查询某计划离线列表
func TaobaoSubwayCampaignOfflineFind(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSubwayCampaignOfflineFindAPIRequest, resp *simba.TaobaoSubwayCampaignOfflineFindAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
