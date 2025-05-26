package baichuan

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/baichuan"
)

// TaobaoBaichuanItemsUnsubscribe 批量删除商品订阅
// taobao.baichuan.items.unsubscribe
//
// 批量删除商品订阅
func TaobaoBaichuanItemsUnsubscribe(ctx context.Context, clt *core.SDKClient, req *baichuan.TaobaoBaichuanItemsUnsubscribeAPIRequest, resp *baichuan.TaobaoBaichuanItemsUnsubscribeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
