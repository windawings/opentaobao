package iot

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/iot"
)

// TaobaoAilabAicloudTopDeviceControlPauseandresume 设备播放暂停
// taobao.ailab.aicloud.top.device.control.pauseandresume
//
// 设备播放暂停
func TaobaoAilabAicloudTopDeviceControlPauseandresume(ctx context.Context, clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopDeviceControlPauseandresumeAPIRequest, resp *iot.TaobaoAilabAicloudTopDeviceControlPauseandresumeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
