package baichuan

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/baichuan"
)

// TaobaoBaichuanTaokeTrace 百川淘客打点
// taobao.baichuan.taoke.trace
//
// 百川淘客打点
func TaobaoBaichuanTaokeTrace(ctx context.Context, clt *core.SDKClient, req *baichuan.TaobaoBaichuanTaokeTraceAPIRequest, resp *baichuan.TaobaoBaichuanTaokeTraceAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
