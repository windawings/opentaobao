package drug

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/drug"
)

// AlibabaHealthNrLogisticsDeliverynoUpdate 上传订单同城快递单号
// alibaba.health.nr.logistics.deliveryno.update
//
// 上传订单同城快递单号
func AlibabaHealthNrLogisticsDeliverynoUpdate(ctx context.Context, clt *core.SDKClient, req *drug.AlibabaHealthNrLogisticsDeliverynoUpdateAPIRequest, resp *drug.AlibabaHealthNrLogisticsDeliverynoUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
