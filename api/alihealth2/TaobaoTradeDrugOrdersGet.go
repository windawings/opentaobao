package alihealth2

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/alihealth2"
)

// TaobaoTradeDrugOrdersGet 阿里健康获取某一药店全部订单
// taobao.trade.drug.orders.get
//
// 阿里健康获取某一药店全部订单
func TaobaoTradeDrugOrdersGet(ctx context.Context, clt *core.SDKClient, req *alihealth2.TaobaoTradeDrugOrdersGetAPIRequest, resp *alihealth2.TaobaoTradeDrugOrdersGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
