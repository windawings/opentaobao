package tmallservice

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tmallservice"
)

// TmallServicecenterWorkcardVerifycodeResend 重发核销码
// tmall.servicecenter.workcard.verifycode.resend
//
// 重发核销码
func TmallServicecenterWorkcardVerifycodeResend(ctx context.Context, clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkcardVerifycodeResendAPIRequest, resp *tmallservice.TmallServicecenterWorkcardVerifycodeResendAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
