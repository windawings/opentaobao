package nazca

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/nazca"
)

// AlibabaNazcaAuthAuthapplyCallback 认证的统一回调接口
// alibaba.nazca.auth.authapply.callback
//
// 认证的统一回调接口
func AlibabaNazcaAuthAuthapplyCallback(ctx context.Context, clt *core.SDKClient, req *nazca.AlibabaNazcaAuthAuthapplyCallbackAPIRequest, resp *nazca.AlibabaNazcaAuthAuthapplyCallbackAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
