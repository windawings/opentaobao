package fans

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/fans"
)

// TmallFansCashpoolCreate 创建资金池
// tmall.fans.cashpool.create
//
// 商家创建资金池接口
func TmallFansCashpoolCreate(ctx context.Context, clt *core.SDKClient, req *fans.TmallFansCashpoolCreateAPIRequest, resp *fans.TmallFansCashpoolCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
