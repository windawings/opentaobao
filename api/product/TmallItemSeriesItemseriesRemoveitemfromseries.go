package product

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/product"
)

// TmallItemSeriesItemseriesRemoveitemfromseries 从商品系列中移除商品
// tmall.item.series.itemseries.removeitemfromseries
//
// 从商品系列中移除商品
func TmallItemSeriesItemseriesRemoveitemfromseries(ctx context.Context, clt *core.SDKClient, req *product.TmallItemSeriesItemseriesRemoveitemfromseriesAPIRequest, resp *product.TmallItemSeriesItemseriesRemoveitemfromseriesAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
