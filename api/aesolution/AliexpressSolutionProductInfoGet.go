package aesolution

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/aesolution"
)

// AliexpressSolutionProductInfoGet Get Single Product Info
// aliexpress.solution.product.info.get
//
// Get Single Product Info
func AliexpressSolutionProductInfoGet(ctx context.Context, clt *core.SDKClient, req *aesolution.AliexpressSolutionProductInfoGetAPIRequest, resp *aesolution.AliexpressSolutionProductInfoGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
