package omniorder

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/omniorder"
)

// AlibabaRetailCommissionOrderQuery 分销订单查询
// alibaba.retail.commission.order.query
//
// 查询商家的分销订单
func AlibabaRetailCommissionOrderQuery(ctx context.Context, clt *core.SDKClient, req *omniorder.AlibabaRetailCommissionOrderQueryAPIRequest, resp *omniorder.AlibabaRetailCommissionOrderQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
