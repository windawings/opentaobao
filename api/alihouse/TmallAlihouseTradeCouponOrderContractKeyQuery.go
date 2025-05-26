package alihouse

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/alihouse"
)

// TmallAlihouseTradeCouponOrderContractKeyQuery 查询电商券履约单合同key
// tmall.alihouse.trade.coupon.order.contract.key.query
//
// 查询电商券履约单合同地址
func TmallAlihouseTradeCouponOrderContractKeyQuery(ctx context.Context, clt *core.SDKClient, req *alihouse.TmallAlihouseTradeCouponOrderContractKeyQueryAPIRequest, resp *alihouse.TmallAlihouseTradeCouponOrderContractKeyQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
