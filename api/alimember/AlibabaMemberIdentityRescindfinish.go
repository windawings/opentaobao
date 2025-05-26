package alimember

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/alimember"
)

// AlibabaMemberIdentityRescindfinish 取消确认
// alibaba.member.identity.rescindfinish
//
// 取消确认
func AlibabaMemberIdentityRescindfinish(ctx context.Context, clt *core.SDKClient, req *alimember.AlibabaMemberIdentityRescindfinishAPIRequest, resp *alimember.AlibabaMemberIdentityRescindfinishAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
