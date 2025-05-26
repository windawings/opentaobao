package charity

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/charity"
)

// AlibabaCharityUserBindGeturi 获取用户绑定uri
// alibaba.charity.user.bind.geturi
//
// 获取用户绑定uri
func AlibabaCharityUserBindGeturi(ctx context.Context, clt *core.SDKClient, req *charity.AlibabaCharityUserBindGeturiAPIRequest, resp *charity.AlibabaCharityUserBindGeturiAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
