package tbitem

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tbitem"
)

// TmallProductSchemaGet 产品信息获取schema获取
// tmall.product.schema.get
//
// 产品信息获取接口schema形式返回
func TmallProductSchemaGet(ctx context.Context, clt *core.SDKClient, req *tbitem.TmallProductSchemaGetAPIRequest, resp *tbitem.TmallProductSchemaGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
