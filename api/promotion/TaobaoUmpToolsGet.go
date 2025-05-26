package promotion

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/promotion"
)

// TaobaoUmpToolsGet 查询工具列表
// taobao.ump.tools.get
//
// 查询工具列表
func TaobaoUmpToolsGet(ctx context.Context, clt *core.SDKClient, req *promotion.TaobaoUmpToolsGetAPIRequest, resp *promotion.TaobaoUmpToolsGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
