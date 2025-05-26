package mozi

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/mozi"
)

// AlibabaMoziBucAccountListAccountids 根据一批账号ID查询账号列表
// alibaba.mozi.buc.account.list.accountids
//
// 根据一批账号ID查询账号列表
func AlibabaMoziBucAccountListAccountids(ctx context.Context, clt *core.SDKClient, req *mozi.AlibabaMoziBucAccountListAccountidsAPIRequest, resp *mozi.AlibabaMoziBucAccountListAccountidsAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
