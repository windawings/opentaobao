package wdk

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/wdk"
)

// AlibabaWdkMarketingItempoolAddcategory 增加商品池里面的类目
// alibaba.wdk.marketing.itempool.addcategory
//
// 增加商品池里面的类目
func AlibabaWdkMarketingItempoolAddcategory(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkMarketingItempoolAddcategoryAPIRequest, resp *wdk.AlibabaWdkMarketingItempoolAddcategoryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
