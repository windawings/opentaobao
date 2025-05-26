package product

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/product"
)

// AlibabaGpuUpdateSchemaGet 获取产品编辑schema规则的接口
// alibaba.gpu.update.schema.get
//
// 获取产品编辑schema规则的接口
func AlibabaGpuUpdateSchemaGet(ctx context.Context, clt *core.SDKClient, req *product.AlibabaGpuUpdateSchemaGetAPIRequest, resp *product.AlibabaGpuUpdateSchemaGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
