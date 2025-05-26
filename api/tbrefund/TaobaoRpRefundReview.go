package tbrefund

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tbrefund"
)

// TaobaoRpRefundReview 审核退款单
// taobao.rp.refund.review
//
// 审核退款单，标志是否可用于批量退款，目前仅支持天猫订单。
func TaobaoRpRefundReview(ctx context.Context, clt *core.SDKClient, req *tbrefund.TaobaoRpRefundReviewAPIRequest, resp *tbrefund.TaobaoRpRefundReviewAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
