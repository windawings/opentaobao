package moscm

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/moscm"
)

// AlibabaMosDeliverySend 发货
// alibaba.mos.delivery.send
//
// 订单发货填写快递单
func AlibabaMosDeliverySend(ctx context.Context, clt *core.SDKClient, req *moscm.AlibabaMosDeliverySendAPIRequest, resp *moscm.AlibabaMosDeliverySendAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
