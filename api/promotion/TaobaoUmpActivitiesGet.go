package promotion

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/promotion"
)

// TaobaoUmpActivitiesGet 查询活动列表
// taobao.ump.activities.get
//
// 查询活动列表
func TaobaoUmpActivitiesGet(ctx context.Context, clt *core.SDKClient, req *promotion.TaobaoUmpActivitiesGetAPIRequest, resp *promotion.TaobaoUmpActivitiesGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
