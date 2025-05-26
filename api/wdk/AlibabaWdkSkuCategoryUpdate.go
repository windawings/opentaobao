package wdk

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/wdk"
)

// AlibabaWdkSkuCategoryUpdate 商家类目修改接口
// alibaba.wdk.sku.category.update
//
// 商家类目修改接口
func AlibabaWdkSkuCategoryUpdate(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkSkuCategoryUpdateAPIRequest, resp *wdk.AlibabaWdkSkuCategoryUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
