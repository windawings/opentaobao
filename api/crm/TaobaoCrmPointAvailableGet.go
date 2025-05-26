package crm

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/crm"
)

// TaobaoCrmPointAvailableGet CRM会员积分查询开放接口
// taobao.crm.point.available.get
//
// 查询用户在某个商家的可用积分数
func TaobaoCrmPointAvailableGet(ctx context.Context, clt *core.SDKClient, req *crm.TaobaoCrmPointAvailableGetAPIRequest, resp *crm.TaobaoCrmPointAvailableGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
