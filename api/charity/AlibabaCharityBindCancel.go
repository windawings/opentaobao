package charity

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/charity"
)

// AlibabaCharityBindCancel 取消用户绑定
// alibaba.charity.bind.cancel
//
// 取消用户绑定
func AlibabaCharityBindCancel(ctx context.Context, clt *core.SDKClient, req *charity.AlibabaCharityBindCancelAPIRequest, resp *charity.AlibabaCharityBindCancelAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
