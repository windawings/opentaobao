package feedflow

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/feedflow"
)

// TaobaoFeedflowItemCrowdModify 覆盖单元下同类型定向人群
// taobao.feedflow.item.crowd.modify
//
// 覆盖单元下同类型定向人群
func TaobaoFeedflowItemCrowdModify(ctx context.Context, clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemCrowdModifyAPIRequest, resp *feedflow.TaobaoFeedflowItemCrowdModifyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
