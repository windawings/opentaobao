package charity

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/charity"
)

// AlibabaValueCoinIssue 爱豆发放
// alibaba.value.coin.issue
//
// 爱豆发放
func AlibabaValueCoinIssue(ctx context.Context, clt *core.SDKClient, req *charity.AlibabaValueCoinIssueAPIRequest, resp *charity.AlibabaValueCoinIssueAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
