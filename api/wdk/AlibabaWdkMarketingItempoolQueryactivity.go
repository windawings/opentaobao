package wdk

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/wdk"
)

// AlibabaWdkMarketingItempoolQueryactivity 查找商品池活动
// alibaba.wdk.marketing.itempool.queryactivity
//
// 查找商品池活动
func AlibabaWdkMarketingItempoolQueryactivity(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkMarketingItempoolQueryactivityAPIRequest, resp *wdk.AlibabaWdkMarketingItempoolQueryactivityAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
