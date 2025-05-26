package feedflow

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/feedflow"
)

// TaobaoFeedflowItemCampaignDelete 删除计划
// taobao.feedflow.item.campaign.delete
//
// 删除计划
func TaobaoFeedflowItemCampaignDelete(ctx context.Context, clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemCampaignDeleteAPIRequest, resp *feedflow.TaobaoFeedflowItemCampaignDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
