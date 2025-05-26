package idle

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/idle"
)

// AlibabaIdleRecycleOrderQuery 闲鱼回收订单查询V1
// alibaba.idle.recycle.order.query
//
// 查询回收订单
func AlibabaIdleRecycleOrderQuery(ctx context.Context, clt *core.SDKClient, req *idle.AlibabaIdleRecycleOrderQueryAPIRequest, resp *idle.AlibabaIdleRecycleOrderQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
