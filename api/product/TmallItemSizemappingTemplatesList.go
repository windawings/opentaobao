package product

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/product"
)

// TmallItemSizemappingTemplatesList 获取天猫商品尺码表模板列表
// tmall.item.sizemapping.templates.list
//
// 获取所有尺码表模板列表。
func TmallItemSizemappingTemplatesList(ctx context.Context, clt *core.SDKClient, req *product.TmallItemSizemappingTemplatesListAPIRequest, resp *product.TmallItemSizemappingTemplatesListAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
