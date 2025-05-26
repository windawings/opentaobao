package aesolution

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/aesolution"
)

// AliexpressSolutionFeedListGet aliexpress.solution.feed.list.get
// aliexpress.solution.feed.list.get
//
// API to query the feed list belonged to a seller
func AliexpressSolutionFeedListGet(ctx context.Context, clt *core.SDKClient, req *aesolution.AliexpressSolutionFeedListGetAPIRequest, resp *aesolution.AliexpressSolutionFeedListGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
