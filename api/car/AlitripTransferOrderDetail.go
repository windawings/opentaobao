package car

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/car"
)

// AlitripTransferOrderDetail 接送订单详情接口
// alitrip.transfer.order.detail
//
// 接送订单详情接口
func AlitripTransferOrderDetail(ctx context.Context, clt *core.SDKClient, req *car.AlitripTransferOrderDetailAPIRequest, resp *car.AlitripTransferOrderDetailAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
