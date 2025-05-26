package user

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/user"
)

// TaobaoOpenAccountCreate Open Account导入数据
// taobao.open.account.create
//
// Open Account导入数据
func TaobaoOpenAccountCreate(ctx context.Context, clt *core.SDKClient, req *user.TaobaoOpenAccountCreateAPIRequest, resp *user.TaobaoOpenAccountCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
