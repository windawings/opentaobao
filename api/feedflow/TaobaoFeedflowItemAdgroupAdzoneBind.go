package feedflow

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/feedflow"
)

// TaobaoFeedflowItemAdgroupAdzoneBind 信息流单元内绑定资源位
// taobao.feedflow.item.adgroup.adzone.bind
//
// 信息流单元内绑定资源位
func TaobaoFeedflowItemAdgroupAdzoneBind(ctx context.Context, clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemAdgroupAdzoneBindAPIRequest, resp *feedflow.TaobaoFeedflowItemAdgroupAdzoneBindAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
