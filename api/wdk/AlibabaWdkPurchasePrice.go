package wdk

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/wdk"
)

// AlibabaWdkPurchasePrice rt回传采购价
// alibaba.wdk.purchase.price
//
// 猫超共享库存项目-rt回传采购价
func AlibabaWdkPurchasePrice(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkPurchasePriceAPIRequest, resp *wdk.AlibabaWdkPurchasePriceAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
