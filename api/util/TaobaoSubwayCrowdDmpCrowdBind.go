package util

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/util"
)

// TaobaoSubwayCrowdDmpCrowdBind 直通车绑定达摩盘人群
// taobao.subway.crowd.dmp.crowd.bind
//
// 直通车绑定达摩盘人群
func TaobaoSubwayCrowdDmpCrowdBind(ctx context.Context, clt *core.SDKClient, req *util.TaobaoSubwayCrowdDmpCrowdBindAPIRequest, resp *util.TaobaoSubwayCrowdDmpCrowdBindAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
