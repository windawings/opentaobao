package wdkitem

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/wdkitem"
)

// AlibabaWdkItemMerchantstoreskuCreate 门店商品信息新建
// alibaba.wdk.item.merchantstoresku.create
//
// 门店商品信息新建
func AlibabaWdkItemMerchantstoreskuCreate(ctx context.Context, clt *core.SDKClient, req *wdkitem.AlibabaWdkItemMerchantstoreskuCreateAPIRequest, resp *wdkitem.AlibabaWdkItemMerchantstoreskuCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
