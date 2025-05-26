package icbushowcase

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/icbushowcase"
)

// AlibabaScbpShowcaseSort 橱窗顺序变更
// alibaba.scbp.showcase.sort
//
// 橱窗顺序变更
func AlibabaScbpShowcaseSort(ctx context.Context, clt *core.SDKClient, req *icbushowcase.AlibabaScbpShowcaseSortAPIRequest, resp *icbushowcase.AlibabaScbpShowcaseSortAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
