package icbu

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/icbu"
)

// AlibabaIcbuCategoryLevelAttrGet 层级属性的子属性获取
// alibaba.icbu.category.level.attr.get
//
// 用于获取层级属性（车型库）的子属性和属性值
func AlibabaIcbuCategoryLevelAttrGet(ctx context.Context, clt *core.SDKClient, req *icbu.AlibabaIcbuCategoryLevelAttrGetAPIRequest, resp *icbu.AlibabaIcbuCategoryLevelAttrGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
