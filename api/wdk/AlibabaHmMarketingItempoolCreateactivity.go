package wdk

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/wdk"
)

// AlibabaHmMarketingItempoolCreateactivity 添加商品池活动
// alibaba.hm.marketing.itempool.createactivity
//
// 添加商品池活动
func AlibabaHmMarketingItempoolCreateactivity(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaHmMarketingItempoolCreateactivityAPIRequest, resp *wdk.AlibabaHmMarketingItempoolCreateactivityAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
