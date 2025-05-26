package aecreatives

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/aecreatives"
)

// AliexpressAffiliateProductSmartmatch 联盟物料智能推荐api
// aliexpress.affiliate.product.smartmatch
//
// 联盟物料算法智能推荐
func AliexpressAffiliateProductSmartmatch(ctx context.Context, clt *core.SDKClient, req *aecreatives.AliexpressAffiliateProductSmartmatchAPIRequest, resp *aecreatives.AliexpressAffiliateProductSmartmatchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
