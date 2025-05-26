package mozi

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/mozi"
)

// AlibabaMoziBucAccountPageall 查询租户内内所有账号
// alibaba.mozi.buc.account.pageall
//
// 查询租户内内所有账号
func AlibabaMoziBucAccountPageall(ctx context.Context, clt *core.SDKClient, req *mozi.AlibabaMoziBucAccountPageallAPIRequest, resp *mozi.AlibabaMoziBucAccountPageallAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
