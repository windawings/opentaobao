package wdk

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/wdk"
)

// AlibabaWdkMarketingFullrangeCreateactivity 创建全场活动
// alibaba.wdk.marketing.fullrange.createactivity
//
// 创建全场活动
func AlibabaWdkMarketingFullrangeCreateactivity(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkMarketingFullrangeCreateactivityAPIRequest, resp *wdk.AlibabaWdkMarketingFullrangeCreateactivityAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
