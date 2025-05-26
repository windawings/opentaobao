package scbp

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/scbp"
)

// AlibabaScbpAdKeywordAdd 外贸直通车加词
// alibaba.scbp.ad.keyword.add
//
// 外贸直通车加词服务
func AlibabaScbpAdKeywordAdd(ctx context.Context, clt *core.SDKClient, req *scbp.AlibabaScbpAdKeywordAddAPIRequest, resp *scbp.AlibabaScbpAdKeywordAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
