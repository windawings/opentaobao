package simba

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/simba"
)

// TaobaoSimbaCreativesChangedGet 分页获取修改过的广告创意ID和修改时间
// taobao.simba.creatives.changed.get
//
// 分页获取修改过的广告创意ID和修改时间
func TaobaoSimbaCreativesChangedGet(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaCreativesChangedGetAPIRequest, resp *simba.TaobaoSimbaCreativesChangedGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
