package simba

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/simba"
)

// TaobaoSimbaCustomersAuthorizedGet 取得当前登录用户的授权账户列表
// taobao.simba.customers.authorized.get
//
// 取得当前登录用户的授权账户列表
func TaobaoSimbaCustomersAuthorizedGet(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaCustomersAuthorizedGetAPIRequest, resp *simba.TaobaoSimbaCustomersAuthorizedGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
