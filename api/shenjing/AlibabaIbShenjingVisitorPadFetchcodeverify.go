package shenjing

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/shenjing"
)

// AlibabaIbShenjingVisitorPadFetchcodeverify 访客通过PAD提交访客码
// alibaba.ib.shenjing.visitor.pad.fetchcodeverify
//
// 访客通过PAD提交访客码，录脸进入园区。
func AlibabaIbShenjingVisitorPadFetchcodeverify(ctx context.Context, clt *core.SDKClient, req *shenjing.AlibabaIbShenjingVisitorPadFetchcodeverifyAPIRequest, resp *shenjing.AlibabaIbShenjingVisitorPadFetchcodeverifyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
