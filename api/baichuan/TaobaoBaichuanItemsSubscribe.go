package baichuan

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/baichuan"
)

// TaobaoBaichuanItemsSubscribe 百川批量商品订阅
// taobao.baichuan.items.subscribe
//
// 百川批量添加订阅的商品
func TaobaoBaichuanItemsSubscribe(ctx context.Context, clt *core.SDKClient, req *baichuan.TaobaoBaichuanItemsSubscribeAPIRequest, resp *baichuan.TaobaoBaichuanItemsSubscribeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
