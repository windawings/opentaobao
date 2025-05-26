package bus

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/bus"
)

// TaobaoBusAgentRefundticketConfirm 商家top回调退款明细
// taobao.bus.agent.refundticket.confirm
//
// 商家通过top回调告知平台退款明细
func TaobaoBusAgentRefundticketConfirm(ctx context.Context, clt *core.SDKClient, req *bus.TaobaoBusAgentRefundticketConfirmAPIRequest, resp *bus.TaobaoBusAgentRefundticketConfirmAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
