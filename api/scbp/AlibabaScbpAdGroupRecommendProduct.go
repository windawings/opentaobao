package scbp

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/scbp"
)

// AlibabaScbpAdGroupRecommendProduct 推品
// alibaba.scbp.ad.group.recommend.product
//
// 推品
func AlibabaScbpAdGroupRecommendProduct(ctx context.Context, clt *core.SDKClient, req *scbp.AlibabaScbpAdGroupRecommendProductAPIRequest, resp *scbp.AlibabaScbpAdGroupRecommendProductAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
