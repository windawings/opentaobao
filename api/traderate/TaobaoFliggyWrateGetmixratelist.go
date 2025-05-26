package traderate

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/traderate"
)

// TaobaoFliggyWrateGetmixratelist 飞猪通用评价接口
// taobao.fliggy.wrate.getmixratelist
//
// 飞猪评价通用接口
func TaobaoFliggyWrateGetmixratelist(ctx context.Context, clt *core.SDKClient, req *traderate.TaobaoFliggyWrateGetmixratelistAPIRequest, resp *traderate.TaobaoFliggyWrateGetmixratelistAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
