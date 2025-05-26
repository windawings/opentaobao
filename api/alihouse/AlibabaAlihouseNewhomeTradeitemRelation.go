package alihouse

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeTradeitemRelation 货独立绑定货品
// alibaba.alihouse.newhome.tradeitem.relation
//
// 货独立绑定货品
func AlibabaAlihouseNewhomeTradeitemRelation(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeTradeitemRelationAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeTradeitemRelationAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
