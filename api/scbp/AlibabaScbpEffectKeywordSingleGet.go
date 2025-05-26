package scbp

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/scbp"
)

// AlibabaScbpEffectKeywordSingleGet 单个关键词效果报表
// alibaba.scbp.effect.keyword.single.get
//
// 单个关键词效果报表
func AlibabaScbpEffectKeywordSingleGet(ctx context.Context, clt *core.SDKClient, req *scbp.AlibabaScbpEffectKeywordSingleGetAPIRequest, resp *scbp.AlibabaScbpEffectKeywordSingleGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
