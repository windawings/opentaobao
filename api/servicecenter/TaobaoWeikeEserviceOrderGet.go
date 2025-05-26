package servicecenter

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/servicecenter"
)

// TaobaoWeikeEserviceOrderGet 客服外包订单查询
// taobao.weike.eservice.order.get
//
// 用于客服外包中服务商查询订单列表
func TaobaoWeikeEserviceOrderGet(ctx context.Context, clt *core.SDKClient, req *servicecenter.TaobaoWeikeEserviceOrderGetAPIRequest, resp *servicecenter.TaobaoWeikeEserviceOrderGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
