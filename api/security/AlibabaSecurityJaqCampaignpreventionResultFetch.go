package security

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/security"
)

// AlibabaSecurityJaqCampaignpreventionResultFetch 获取活动保护结果
// alibaba.security.jaq.campaignprevention.result.fetch
//
// 获取活动保护结果
func AlibabaSecurityJaqCampaignpreventionResultFetch(ctx context.Context, clt *core.SDKClient, req *security.AlibabaSecurityJaqCampaignpreventionResultFetchAPIRequest, resp *security.AlibabaSecurityJaqCampaignpreventionResultFetchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
