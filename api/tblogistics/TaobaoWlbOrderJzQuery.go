package tblogistics

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tblogistics"
)

// TaobaoWlbOrderJzQuery 家装业务查询物流公司api
// taobao.wlb.order.jz.query
//
// 家装业务查询物流公司api
func TaobaoWlbOrderJzQuery(ctx context.Context, clt *core.SDKClient, req *tblogistics.TaobaoWlbOrderJzQueryAPIRequest, resp *tblogistics.TaobaoWlbOrderJzQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
