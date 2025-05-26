package interact

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/interact"
)

// AlibabaInteractSensorAuthorize 客户端授权页
// alibaba.interact.sensor.authorize
//
// 客户端授权页
func AlibabaInteractSensorAuthorize(ctx context.Context, clt *core.SDKClient, req *interact.AlibabaInteractSensorAuthorizeAPIRequest, resp *interact.AlibabaInteractSensorAuthorizeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
