package aeusergrowth

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/aeusergrowth"
)

// AliexpressUsergrowthRecommendItemsGet 第三方平台推荐商品
// aliexpress.usergrowth.recommend.items.get
//
// 第三方平台的推荐AE商品   场景：skin 、底部推荐等
func AliexpressUsergrowthRecommendItemsGet(ctx context.Context, clt *core.SDKClient, req *aeusergrowth.AliexpressUsergrowthRecommendItemsGetAPIRequest, resp *aeusergrowth.AliexpressUsergrowthRecommendItemsGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
