package interact

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/interact"
)

// AlibabaInteractSensorGyro 陀螺仪
// alibaba.interact.sensor.gyro
//
// 客户端陀螺仪
func AlibabaInteractSensorGyro(ctx context.Context, clt *core.SDKClient, req *interact.AlibabaInteractSensorGyroAPIRequest, resp *interact.AlibabaInteractSensorGyroAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
