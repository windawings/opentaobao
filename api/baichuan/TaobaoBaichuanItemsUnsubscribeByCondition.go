package baichuan

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/baichuan"
)

// TaobaoBaichuanItemsUnsubscribeByCondition 根据条件删除订阅关系
// taobao.baichuan.items.unsubscribe.by.condition
//
// 根据条件删除订阅关系
func TaobaoBaichuanItemsUnsubscribeByCondition(ctx context.Context, clt *core.SDKClient, req *baichuan.TaobaoBaichuanItemsUnsubscribeByConditionAPIRequest, resp *baichuan.TaobaoBaichuanItemsUnsubscribeByConditionAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
