package tbitem

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tbitem"
)

// TmallItemCombineGet 组合商品获取接口
// tmall.item.combine.get
//
// 查询组合商品的SKU信息
func TmallItemCombineGet(ctx context.Context, clt *core.SDKClient, req *tbitem.TmallItemCombineGetAPIRequest, resp *tbitem.TmallItemCombineGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
