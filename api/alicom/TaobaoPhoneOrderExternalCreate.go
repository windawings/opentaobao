package alicom

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/alicom"
)

// TaobaoPhoneOrderExternalCreate 数字虚拟话费外放下单接口
// taobao.phone.order.external.create
//
// 数字虚拟话费外放下单接口
func TaobaoPhoneOrderExternalCreate(ctx context.Context, clt *core.SDKClient, req *alicom.TaobaoPhoneOrderExternalCreateAPIRequest, resp *alicom.TaobaoPhoneOrderExternalCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
