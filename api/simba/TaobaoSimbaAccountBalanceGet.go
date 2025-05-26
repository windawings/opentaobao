package simba

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/simba"
)

// TaobaoSimbaAccountBalanceGet 获取实时余额，”元”为单位
// taobao.simba.account.balance.get
//
// 获取实时余额，”元”为单位
func TaobaoSimbaAccountBalanceGet(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaAccountBalanceGetAPIRequest, resp *simba.TaobaoSimbaAccountBalanceGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
