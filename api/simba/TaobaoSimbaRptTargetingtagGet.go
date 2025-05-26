package simba

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/simba"
)

// TaobaoSimbaRptTargetingtagGet 搜索人群离线报表
// taobao.simba.rpt.targetingtag.get
//
// 获取搜搜人群实时报表
func TaobaoSimbaRptTargetingtagGet(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaRptTargetingtagGetAPIRequest, resp *simba.TaobaoSimbaRptTargetingtagGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
