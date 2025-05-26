package feedflow

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/feedflow"
)

// TaobaoFeedflowItemAdgroupModify 信息流单元修改
// taobao.feedflow.item.adgroup.modify
//
// 信息流单元修改
func TaobaoFeedflowItemAdgroupModify(ctx context.Context, clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemAdgroupModifyAPIRequest, resp *feedflow.TaobaoFeedflowItemAdgroupModifyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
