package alicom

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/alicom"
)

// TaobaoPhoneItemExternalRecommend 话费选品能力外放
// taobao.phone.item.external.recommend
//
// 话费选品能力外放
func TaobaoPhoneItemExternalRecommend(ctx context.Context, clt *core.SDKClient, req *alicom.TaobaoPhoneItemExternalRecommendAPIRequest, resp *alicom.TaobaoPhoneItemExternalRecommendAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
