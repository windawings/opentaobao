package scbp

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/scbp"
)

// AlibabaScbpAdKeywordTagUpdate 修改关键词所属分组
// alibaba.scbp.ad.keyword.tag.update
//
// 修改关键词所属分组
func AlibabaScbpAdKeywordTagUpdate(ctx context.Context, clt *core.SDKClient, req *scbp.AlibabaScbpAdKeywordTagUpdateAPIRequest, resp *scbp.AlibabaScbpAdKeywordTagUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
