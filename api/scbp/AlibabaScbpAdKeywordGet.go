package scbp

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/scbp"
)

// AlibabaScbpAdKeywordGet 外贸直通车查询关键词
// alibaba.scbp.ad.keyword.get
//
// 外贸直通车查询关键词
func AlibabaScbpAdKeywordGet(ctx context.Context, clt *core.SDKClient, req *scbp.AlibabaScbpAdKeywordGetAPIRequest, resp *scbp.AlibabaScbpAdKeywordGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
