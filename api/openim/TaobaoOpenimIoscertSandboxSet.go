package openim

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/openim"
)

// TaobaoOpenimIoscertSandboxSet 设置开发环境证书
// taobao.openim.ioscert.sandbox.set
//
// 设置开发环境证书
func TaobaoOpenimIoscertSandboxSet(ctx context.Context, clt *core.SDKClient, req *openim.TaobaoOpenimIoscertSandboxSetAPIRequest, resp *openim.TaobaoOpenimIoscertSandboxSetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
