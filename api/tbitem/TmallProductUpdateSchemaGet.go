package tbitem

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tbitem"
)

// TmallProductUpdateSchemaGet 产品更新规则获取接口
// tmall.product.update.schema.get
//
// 获取用户更新产品的规则
func TmallProductUpdateSchemaGet(ctx context.Context, clt *core.SDKClient, req *tbitem.TmallProductUpdateSchemaGetAPIRequest, resp *tbitem.TmallProductUpdateSchemaGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
