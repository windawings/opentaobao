package tblogistics

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tblogistics"
)

// TaobaoWlbOrderJzConsign 家装发货接口
// taobao.wlb.order.jz.consign
//
// 家装类订单使用该接口发货
func TaobaoWlbOrderJzConsign(ctx context.Context, clt *core.SDKClient, req *tblogistics.TaobaoWlbOrderJzConsignAPIRequest, resp *tblogistics.TaobaoWlbOrderJzConsignAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
