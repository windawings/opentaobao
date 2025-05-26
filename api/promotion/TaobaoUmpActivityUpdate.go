package promotion

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/promotion"
)

// TaobaoUmpActivityUpdate 修改活动信息
// taobao.ump.activity.update
//
// 修改营销活动
func TaobaoUmpActivityUpdate(ctx context.Context, clt *core.SDKClient, req *promotion.TaobaoUmpActivityUpdateAPIRequest, resp *promotion.TaobaoUmpActivityUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
