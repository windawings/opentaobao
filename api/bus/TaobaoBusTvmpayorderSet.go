package bus

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/bus"
)

// TaobaoBusTvmpayorderSet 自助机条形码被动支付
// taobao.bus.tvmpayorder.set
//
// 汽车票线下自助机条形码支付
func TaobaoBusTvmpayorderSet(ctx context.Context, clt *core.SDKClient, req *bus.TaobaoBusTvmpayorderSetAPIRequest, resp *bus.TaobaoBusTvmpayorderSetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
