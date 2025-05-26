package aesolution

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/aesolution"
)

// AliexpressSolutionBatchProductPriceUpdate aliexpress.solution.batch.product.price.update
// aliexpress.solution.batch.product.price.update
//
// batch product price update operation for oversea sellers
func AliexpressSolutionBatchProductPriceUpdate(ctx context.Context, clt *core.SDKClient, req *aesolution.AliexpressSolutionBatchProductPriceUpdateAPIRequest, resp *aesolution.AliexpressSolutionBatchProductPriceUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
