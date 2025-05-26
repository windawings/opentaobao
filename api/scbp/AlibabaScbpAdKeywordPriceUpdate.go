package scbp

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/scbp"
)

// AlibabaScbpAdKeywordPriceUpdate 关键词改价
// alibaba.scbp.ad.keyword.price.update
//
// 关键词改价
func AlibabaScbpAdKeywordPriceUpdate(ctx context.Context, clt *core.SDKClient, req *scbp.AlibabaScbpAdKeywordPriceUpdateAPIRequest, resp *scbp.AlibabaScbpAdKeywordPriceUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
