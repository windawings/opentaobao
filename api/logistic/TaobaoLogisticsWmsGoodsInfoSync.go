package logistic

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/logistic"
)

// TaobaoLogisticsWmsGoodsInfoSync WMS回传货品长宽高图片等信息
// taobao.logistics.wms.goods.info.sync
//
// WMS回传货品长宽高图片等信息
func TaobaoLogisticsWmsGoodsInfoSync(ctx context.Context, clt *core.SDKClient, req *logistic.TaobaoLogisticsWmsGoodsInfoSyncAPIRequest, resp *logistic.TaobaoLogisticsWmsGoodsInfoSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
