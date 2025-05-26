package scbp

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/scbp"
)

// AlibabaScbpAdKeywordStatusUpdate 关键词启动暂停推广
// alibaba.scbp.ad.keyword.status.update
//
// 关键词启动暂停推广
func AlibabaScbpAdKeywordStatusUpdate(ctx context.Context, clt *core.SDKClient, req *scbp.AlibabaScbpAdKeywordStatusUpdateAPIRequest, resp *scbp.AlibabaScbpAdKeywordStatusUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
