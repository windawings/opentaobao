package iot

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/iot"
)

// TaobaoAilabAicloudTopDeviceControlPlaybyid 通过id播放歌曲
// taobao.ailab.aicloud.top.device.control.playbyid
//
// 通过id播放歌曲
func TaobaoAilabAicloudTopDeviceControlPlaybyid(ctx context.Context, clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopDeviceControlPlaybyidAPIRequest, resp *iot.TaobaoAilabAicloudTopDeviceControlPlaybyidAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
