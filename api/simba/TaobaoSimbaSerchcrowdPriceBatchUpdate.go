package simba

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/simba"
)

// TaobaoSimbaSerchcrowdPriceBatchUpdate 单品推广搜索人群修改溢价
// taobao.simba.serchcrowd.price.batch.update
//
// 单品推广搜索人群修改溢价, 不支持跨推广单元修改
func TaobaoSimbaSerchcrowdPriceBatchUpdate(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaSerchcrowdPriceBatchUpdateAPIRequest, resp *simba.TaobaoSimbaSerchcrowdPriceBatchUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
