package damai

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/damai"
)

// AlibabaDamaiMevOpenPushperform 大麦换验平台-第三方对外开放-场次接口pushPerform
// alibaba.damai.mev.open.pushperform
//
// pushPerform
func AlibabaDamaiMevOpenPushperform(ctx context.Context, clt *core.SDKClient, req *damai.AlibabaDamaiMevOpenPushperformAPIRequest, resp *damai.AlibabaDamaiMevOpenPushperformAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
