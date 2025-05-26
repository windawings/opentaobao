package feedflow

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/feedflow"
)

// TaobaoFeedflowItemAdgroupDelete 根据单元id删除单元
// taobao.feedflow.item.adgroup.delete
//
// 根据单元id删除单元
func TaobaoFeedflowItemAdgroupDelete(ctx context.Context, clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemAdgroupDeleteAPIRequest, resp *feedflow.TaobaoFeedflowItemAdgroupDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
