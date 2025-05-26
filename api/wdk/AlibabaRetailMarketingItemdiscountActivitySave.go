package wdk

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/wdk"
)

// AlibabaRetailMarketingItemdiscountActivitySave 【同城零售】单品活动保存
// alibaba.retail.marketing.itemdiscount.activity.save
//
// 【同城零售】单品活动保存
func AlibabaRetailMarketingItemdiscountActivitySave(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaRetailMarketingItemdiscountActivitySaveAPIRequest, resp *wdk.AlibabaRetailMarketingItemdiscountActivitySaveAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
