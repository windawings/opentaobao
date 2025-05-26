package iot

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/iot"
)

// TaobaoAilabAicloudTopDeviceControlHibernation 定时休眠
// taobao.ailab.aicloud.top.device.control.hibernation
//
// 定时休眠
func TaobaoAilabAicloudTopDeviceControlHibernation(ctx context.Context, clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopDeviceControlHibernationAPIRequest, resp *iot.TaobaoAilabAicloudTopDeviceControlHibernationAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
