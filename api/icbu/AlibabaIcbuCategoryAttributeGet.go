package icbu

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/icbu"
)

// AlibabaIcbuCategoryAttributeGet 类目属性获取
// alibaba.icbu.category.attribute.get
//
// 根据类目ID获取系统定义的属性
func AlibabaIcbuCategoryAttributeGet(ctx context.Context, clt *core.SDKClient, req *icbu.AlibabaIcbuCategoryAttributeGetAPIRequest, resp *icbu.AlibabaIcbuCategoryAttributeGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
