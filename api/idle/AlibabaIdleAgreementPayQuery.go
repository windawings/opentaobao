package idle

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/idle"
)

// AlibabaIdleAgreementPayQuery 代扣详情查询
// alibaba.idle.agreement.pay.query
//
// 查询代扣结果
func AlibabaIdleAgreementPayQuery(ctx context.Context, clt *core.SDKClient, req *idle.AlibabaIdleAgreementPayQueryAPIRequest, resp *idle.AlibabaIdleAgreementPayQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
