package alicom

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/alicom"
)

// TaobaoPhoneOrderExternalStatus 话费外放订单状态接口
// taobao.phone.order.external.status
//
// 话费外放订单状态接口
func TaobaoPhoneOrderExternalStatus(ctx context.Context, clt *core.SDKClient, req *alicom.TaobaoPhoneOrderExternalStatusAPIRequest, resp *alicom.TaobaoPhoneOrderExternalStatusAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
