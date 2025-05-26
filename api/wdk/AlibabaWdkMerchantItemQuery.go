package wdk

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/wdk"
)

// AlibabaWdkMerchantItemQuery 商家商品查询
// alibaba.wdk.merchant.item.query
//
// 商家商品查询
func AlibabaWdkMerchantItemQuery(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkMerchantItemQueryAPIRequest, resp *wdk.AlibabaWdkMerchantItemQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
