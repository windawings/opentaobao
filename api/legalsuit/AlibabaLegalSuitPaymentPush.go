package legalsuit

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/legalsuit"
)

// AlibabaLegalSuitPaymentPush 外部推送缴费
// alibaba.legal.suit.payment.push
//
// 外部推送缴费
func AlibabaLegalSuitPaymentPush(ctx context.Context, clt *core.SDKClient, req *legalsuit.AlibabaLegalSuitPaymentPushAPIRequest, resp *legalsuit.AlibabaLegalSuitPaymentPushAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
