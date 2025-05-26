package ju

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/ju"
)

// AlibabaJhsCommunitySubmittingText 聚划算社群动态文案下发接口
// alibaba.jhs.community.submitting.text
//
// 聚划算社群动态文案下发接口
func AlibabaJhsCommunitySubmittingText(ctx context.Context, clt *core.SDKClient, req *ju.AlibabaJhsCommunitySubmittingTextAPIRequest, resp *ju.AlibabaJhsCommunitySubmittingTextAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
