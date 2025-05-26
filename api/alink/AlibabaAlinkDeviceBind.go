package alink

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/alink"
)

// AlibabaAlinkDeviceBind 绑定设备
// alibaba.alink.device.bind
//
// 阿里智能解绑设备
func AlibabaAlinkDeviceBind(ctx context.Context, clt *core.SDKClient, req *alink.AlibabaAlinkDeviceBindAPIRequest, resp *alink.AlibabaAlinkDeviceBindAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
