package scbp

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/scbp"
)

// AlibabaScbpAdCampaignDeleteForbiddenKeyword 删除屏蔽词
// alibaba.scbp.ad.campaign.delete.forbidden.keyword
//
// 删除屏蔽词
func AlibabaScbpAdCampaignDeleteForbiddenKeyword(ctx context.Context, clt *core.SDKClient, req *scbp.AlibabaScbpAdCampaignDeleteForbiddenKeywordAPIRequest, resp *scbp.AlibabaScbpAdCampaignDeleteForbiddenKeywordAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
