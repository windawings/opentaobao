package charity

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/charity"
)

// AlibabaValueUserGetid 获取用户userId
// alibaba.value.user.getid
//
// 获取用户userId
func AlibabaValueUserGetid(ctx context.Context, clt *core.SDKClient, req *charity.AlibabaValueUserGetidAPIRequest, resp *charity.AlibabaValueUserGetidAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
