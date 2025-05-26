package campus

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/campus"
)

// AlibabaCampusGuardantGatewayCallback 人卡关系回调
// alibaba.campus.guardant.gateway.callback
//
// 门禁供应商回调平台通知同步结果
func AlibabaCampusGuardantGatewayCallback(ctx context.Context, clt *core.SDKClient, req *campus.AlibabaCampusGuardantGatewayCallbackAPIRequest, resp *campus.AlibabaCampusGuardantGatewayCallbackAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
