package simba

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/simba"
)

// TaobaoSimbaRptTargetingtageffectGet 获取定向效果报表数据
// taobao.simba.rpt.targetingtageffect.get
//
// 获取定向效果报表数据
func TaobaoSimbaRptTargetingtageffectGet(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaRptTargetingtageffectGetAPIRequest, resp *simba.TaobaoSimbaRptTargetingtageffectGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
