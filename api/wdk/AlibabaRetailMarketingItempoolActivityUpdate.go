package wdk

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/wdk"
)

// AlibabaRetailMarketingItempoolActivityUpdate 更新商品池活动【同城零售】
// alibaba.retail.marketing.itempool.activity.update
//
// 同城零售商品池活动更新
func AlibabaRetailMarketingItempoolActivityUpdate(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaRetailMarketingItempoolActivityUpdateAPIRequest, resp *wdk.AlibabaRetailMarketingItempoolActivityUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
