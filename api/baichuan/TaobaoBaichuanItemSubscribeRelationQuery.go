package baichuan

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/baichuan"
)

// TaobaoBaichuanItemSubscribeRelationQuery 查询单个订阅关系
// taobao.baichuan.item.subscribe.relation.query
//
// 查询单个订阅关系
func TaobaoBaichuanItemSubscribeRelationQuery(ctx context.Context, clt *core.SDKClient, req *baichuan.TaobaoBaichuanItemSubscribeRelationQueryAPIRequest, resp *baichuan.TaobaoBaichuanItemSubscribeRelationQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
