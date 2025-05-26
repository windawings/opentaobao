package user

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/user"
)

// TaobaoOpenAccountUpdate Open Account数据更新
// taobao.open.account.update
//
// Open Account数据更新
func TaobaoOpenAccountUpdate(ctx context.Context, clt *core.SDKClient, req *user.TaobaoOpenAccountUpdateAPIRequest, resp *user.TaobaoOpenAccountUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
