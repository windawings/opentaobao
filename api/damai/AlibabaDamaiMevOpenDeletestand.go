package damai

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/damai"
)

// AlibabaDamaiMevOpenDeletestand 大麦换验平台-第三方对外开放-看台接口deleteStand
// alibaba.damai.mev.open.deletestand
//
// deleteStand
func AlibabaDamaiMevOpenDeletestand(ctx context.Context, clt *core.SDKClient, req *damai.AlibabaDamaiMevOpenDeletestandAPIRequest, resp *damai.AlibabaDamaiMevOpenDeletestandAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
