package wdk

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/wdk"
)

// AlibabaRetailMarketingItemdiscountActivitySkuDelete 删除特价活动商品
// alibaba.retail.marketing.itemdiscount.activity.sku.delete
//
// 删除活动商品信息【同城零售】
func AlibabaRetailMarketingItemdiscountActivitySkuDelete(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaRetailMarketingItemdiscountActivitySkuDeleteAPIRequest, resp *wdk.AlibabaRetailMarketingItemdiscountActivitySkuDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
