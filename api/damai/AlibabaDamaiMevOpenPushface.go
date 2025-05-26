package damai

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/damai"
)

// AlibabaDamaiMevOpenPushface 大麦换验平台-第三方对外开放-票面接口pushFace
// alibaba.damai.mev.open.pushface
//
// pushFace
func AlibabaDamaiMevOpenPushface(ctx context.Context, clt *core.SDKClient, req *damai.AlibabaDamaiMevOpenPushfaceAPIRequest, resp *damai.AlibabaDamaiMevOpenPushfaceAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
