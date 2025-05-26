package feedflow

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/feedflow"
)

// TaobaoFeedflowItemCampaignGet 通过计划id查询计划
// taobao.feedflow.item.campaign.get
//
// 通过计划id查询计划
func TaobaoFeedflowItemCampaignGet(ctx context.Context, clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemCampaignGetAPIRequest, resp *feedflow.TaobaoFeedflowItemCampaignGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
