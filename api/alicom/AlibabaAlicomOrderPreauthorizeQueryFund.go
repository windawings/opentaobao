package alicom

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/alicom"
)

// AlibabaAlicomOrderPreauthorizeQueryFund 资金流水查询
// alibaba.alicom.order.preauthorize.query.fund
//
// 预授权-资金流水查询
func AlibabaAlicomOrderPreauthorizeQueryFund(ctx context.Context, clt *core.SDKClient, req *alicom.AlibabaAlicomOrderPreauthorizeQueryFundAPIRequest, resp *alicom.AlibabaAlicomOrderPreauthorizeQueryFundAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
