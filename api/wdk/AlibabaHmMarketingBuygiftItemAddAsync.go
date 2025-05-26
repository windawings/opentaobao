package wdk

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/wdk"
)

// AlibabaHmMarketingBuygiftItemAddAsync 批量发布买赠商品
// alibaba.hm.marketing.buygift.item.add.async
//
// 批量发布买赠商品
func AlibabaHmMarketingBuygiftItemAddAsync(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaHmMarketingBuygiftItemAddAsyncAPIRequest, resp *wdk.AlibabaHmMarketingBuygiftItemAddAsyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
