package product

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/product"
)

// AlibabaGpuAddSchemaGet 获取产品发布规则接口
// alibaba.gpu.add.schema.get
//
// 获取产品发布规则接口
func AlibabaGpuAddSchemaGet(ctx context.Context, clt *core.SDKClient, req *product.AlibabaGpuAddSchemaGetAPIRequest, resp *product.AlibabaGpuAddSchemaGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
