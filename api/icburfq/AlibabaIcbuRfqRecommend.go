package icburfq

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/icburfq"
)

// AlibabaIcbuRfqRecommend rfq推荐
// alibaba.icbu.rfq.recommend
//
// rfq推荐
func AlibabaIcbuRfqRecommend(ctx context.Context, clt *core.SDKClient, req *icburfq.AlibabaIcbuRfqRecommendAPIRequest, resp *icburfq.AlibabaIcbuRfqRecommendAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
