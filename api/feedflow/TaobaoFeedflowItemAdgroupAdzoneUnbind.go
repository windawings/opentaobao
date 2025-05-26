package feedflow

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/feedflow"
)

// TaobaoFeedflowItemAdgroupAdzoneUnbind 信息流单元内解绑资源位
// taobao.feedflow.item.adgroup.adzone.unbind
//
// 信息流单元内解绑资源位
func TaobaoFeedflowItemAdgroupAdzoneUnbind(ctx context.Context, clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemAdgroupAdzoneUnbindAPIRequest, resp *feedflow.TaobaoFeedflowItemAdgroupAdzoneUnbindAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
