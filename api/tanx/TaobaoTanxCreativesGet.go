package tanx

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tanx"
)

// TaobaoTanxCreativesGet 批量获取DSP用户的创意审核结果
// taobao.tanx.creatives.get
//
// 批量获取DSP用户的创意审核结果
func TaobaoTanxCreativesGet(ctx context.Context, clt *core.SDKClient, req *tanx.TaobaoTanxCreativesGetAPIRequest, resp *tanx.TaobaoTanxCreativesGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
