package interact

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/interact"
)

// AlibabaInteractSensorShake 摇一摇
// alibaba.interact.sensor.shake
//
// 摇一摇
func AlibabaInteractSensorShake(ctx context.Context, clt *core.SDKClient, req *interact.AlibabaInteractSensorShakeAPIRequest, resp *interact.AlibabaInteractSensorShakeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
