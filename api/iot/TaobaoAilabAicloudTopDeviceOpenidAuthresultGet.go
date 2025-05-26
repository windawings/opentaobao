package iot

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/iot"
)

// TaobaoAilabAicloudTopDeviceOpenidAuthresultGet 获取openId设备授权码验证结果
// taobao.ailab.aicloud.top.device.openid.authresult.get
//
// 获取openId设备授权码验证结果
func TaobaoAilabAicloudTopDeviceOpenidAuthresultGet(ctx context.Context, clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopDeviceOpenidAuthresultGetAPIRequest, resp *iot.TaobaoAilabAicloudTopDeviceOpenidAuthresultGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
