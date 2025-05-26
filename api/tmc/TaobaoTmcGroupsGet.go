package tmc

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tmc"
)

// TaobaoTmcGroupsGet 获取自定义用户分组列表
// taobao.tmc.groups.get
//
// 获取自定义用户分组列表
func TaobaoTmcGroupsGet(ctx context.Context, clt *core.SDKClient, req *tmc.TaobaoTmcGroupsGetAPIRequest, resp *tmc.TaobaoTmcGroupsGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
