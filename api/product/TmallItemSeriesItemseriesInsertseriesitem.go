package product

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/product"
)

// TmallItemSeriesItemseriesInsertseriesitem 向系列中添加系列商品
// tmall.item.series.itemseries.insertseriesitem
//
// 向系列中添加系列商品
func TmallItemSeriesItemseriesInsertseriesitem(ctx context.Context, clt *core.SDKClient, req *product.TmallItemSeriesItemseriesInsertseriesitemAPIRequest, resp *product.TmallItemSeriesItemseriesInsertseriesitemAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
