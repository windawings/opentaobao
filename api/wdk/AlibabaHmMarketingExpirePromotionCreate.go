package wdk

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/wdk"
)

// AlibabaHmMarketingExpirePromotionCreate 短保优惠创建
// alibaba.hm.marketing.expire.promotion.create
//
// 过期优惠优惠信息录入
func AlibabaHmMarketingExpirePromotionCreate(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaHmMarketingExpirePromotionCreateAPIRequest, resp *wdk.AlibabaHmMarketingExpirePromotionCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
