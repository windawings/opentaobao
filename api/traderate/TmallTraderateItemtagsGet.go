package traderate

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/traderate"
)

// TmallTraderateItemtagsGet 通过商品ID获取标签列表
// tmall.traderate.itemtags.get
//
// 通过商品ID获取标签详细信息
func TmallTraderateItemtagsGet(ctx context.Context, clt *core.SDKClient, req *traderate.TmallTraderateItemtagsGetAPIRequest, resp *traderate.TmallTraderateItemtagsGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
