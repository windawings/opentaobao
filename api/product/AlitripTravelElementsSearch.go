package product

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/product"
)

// AlitripTravelElementsSearch 商家元素搜索
// alitrip.travel.elements.search
//
// 提供商家维护的景点、酒店、餐饮等元素搜索
func AlitripTravelElementsSearch(ctx context.Context, clt *core.SDKClient, req *product.AlitripTravelElementsSearchAPIRequest, resp *product.AlitripTravelElementsSearchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
