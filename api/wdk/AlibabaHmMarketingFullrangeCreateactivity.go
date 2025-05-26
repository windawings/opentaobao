package wdk

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/wdk"
)

// AlibabaHmMarketingFullrangeCreateactivity 创建全场活动
// alibaba.hm.marketing.fullrange.createactivity
//
// 创建全场活动
func AlibabaHmMarketingFullrangeCreateactivity(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaHmMarketingFullrangeCreateactivityAPIRequest, resp *wdk.AlibabaHmMarketingFullrangeCreateactivityAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
