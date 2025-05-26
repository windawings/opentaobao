package elife

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/elife"
)

// TaobaoElifeLifecardRefund 品牌惠卡券冲正退还
// taobao.elife.lifecard.refund
//
// 淘宝生活汇消费卡虚拟卡，线下冲正退货接口
func TaobaoElifeLifecardRefund(ctx context.Context, clt *core.SDKClient, req *elife.TaobaoElifeLifecardRefundAPIRequest, resp *elife.TaobaoElifeLifecardRefundAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
