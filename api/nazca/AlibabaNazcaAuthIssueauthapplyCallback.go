package nazca

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/nazca"
)

// AlibabaNazcaAuthIssueauthapplyCallback 出证申请回调
// alibaba.nazca.auth.issueauthapply.callback
//
// 出证申请回调
func AlibabaNazcaAuthIssueauthapplyCallback(ctx context.Context, clt *core.SDKClient, req *nazca.AlibabaNazcaAuthIssueauthapplyCallbackAPIRequest, resp *nazca.AlibabaNazcaAuthIssueauthapplyCallbackAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
