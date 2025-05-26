package wdk

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/wdk"
)

// AlibabaHmMarketingItemdiscountQueryitems 查询特价商品
// alibaba.hm.marketing.itemdiscount.queryitems
//
// 查询参加特价活动的商品优惠详情
func AlibabaHmMarketingItemdiscountQueryitems(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaHmMarketingItemdiscountQueryitemsAPIRequest, resp *wdk.AlibabaHmMarketingItemdiscountQueryitemsAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
