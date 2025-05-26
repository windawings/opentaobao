package alihealth2

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/alihealth2"
)

// TaobaoTradeDrugRefuseorder 阿里健康020拒单
// taobao.trade.drug.refuseorder
//
// 阿里健康020拒单
func TaobaoTradeDrugRefuseorder(ctx context.Context, clt *core.SDKClient, req *alihealth2.TaobaoTradeDrugRefuseorderAPIRequest, resp *alihealth2.TaobaoTradeDrugRefuseorderAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
