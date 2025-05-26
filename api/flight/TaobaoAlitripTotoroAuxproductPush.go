package flight

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/flight"
)

// TaobaoAlitripTotoroAuxproductPush 廉航辅营产品投放
// taobao.alitrip.totoro.auxproduct.push
//
// 廉航辅营产品投放接口
func TaobaoAlitripTotoroAuxproductPush(ctx context.Context, clt *core.SDKClient, req *flight.TaobaoAlitripTotoroAuxproductPushAPIRequest, resp *flight.TaobaoAlitripTotoroAuxproductPushAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
