package nazca

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/nazca"
)

// AlibabaNazcaAuthChangeauthapplyCallback 变更认证回调
// alibaba.nazca.auth.changeauthapply.callback
//
// 变更认证回调
func AlibabaNazcaAuthChangeauthapplyCallback(ctx context.Context, clt *core.SDKClient, req *nazca.AlibabaNazcaAuthChangeauthapplyCallbackAPIRequest, resp *nazca.AlibabaNazcaAuthChangeauthapplyCallbackAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
