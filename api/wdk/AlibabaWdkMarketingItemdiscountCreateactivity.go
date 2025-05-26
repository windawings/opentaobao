package wdk

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/wdk"
)

// AlibabaWdkMarketingItemdiscountCreateactivity 创建商品特价活动
// alibaba.wdk.marketing.itemdiscount.createactivity
//
// 创建商品特价活动
func AlibabaWdkMarketingItemdiscountCreateactivity(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkMarketingItemdiscountCreateactivityAPIRequest, resp *wdk.AlibabaWdkMarketingItemdiscountCreateactivityAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
