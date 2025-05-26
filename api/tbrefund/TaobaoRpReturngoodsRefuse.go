package tbrefund

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tbrefund"
)

// TaobaoRpReturngoodsRefuse 卖家拒绝退货
// taobao.rp.returngoods.refuse
//
// 卖家拒绝退货，目前仅支持天猫退货。
func TaobaoRpReturngoodsRefuse(ctx context.Context, clt *core.SDKClient, req *tbrefund.TaobaoRpReturngoodsRefuseAPIRequest, resp *tbrefund.TaobaoRpReturngoodsRefuseAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
