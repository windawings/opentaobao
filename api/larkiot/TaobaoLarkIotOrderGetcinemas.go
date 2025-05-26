package larkiot

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/larkiot"
)

// TaobaoLarkIotOrderGetcinemas 获取iot渠道开放的影院
// taobao.lark.iot.order.getcinemas
//
// iot渠道拉取有权限访问的影院
func TaobaoLarkIotOrderGetcinemas(ctx context.Context, clt *core.SDKClient, req *larkiot.TaobaoLarkIotOrderGetcinemasAPIRequest, resp *larkiot.TaobaoLarkIotOrderGetcinemasAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
