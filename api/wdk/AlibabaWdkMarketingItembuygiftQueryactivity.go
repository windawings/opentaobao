package wdk

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/wdk"
)

// AlibabaWdkMarketingItembuygiftQueryactivity 查询买赠活动
// alibaba.wdk.marketing.itembuygift.queryactivity
//
// 查询买赠活动
func AlibabaWdkMarketingItembuygiftQueryactivity(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkMarketingItembuygiftQueryactivityAPIRequest, resp *wdk.AlibabaWdkMarketingItembuygiftQueryactivityAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
