package wdk

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/wdk"
)

// AlibabaHmMarketingItempoolStairAdditem 商品池阶梯商品添加
// alibaba.hm.marketing.itempool.stair.additem
//
// 添加商品池阶梯商品
func AlibabaHmMarketingItempoolStairAdditem(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaHmMarketingItempoolStairAdditemAPIRequest, resp *wdk.AlibabaHmMarketingItempoolStairAdditemAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
