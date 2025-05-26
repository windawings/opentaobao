package idle

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/idle"
)

// AlibabaIdleTenderAftersaleOrderGet 闲鱼帮卖售后服务单查询
// alibaba.idle.tender.aftersale.order.get
//
// 闲鱼帮卖售后服务单查询
func AlibabaIdleTenderAftersaleOrderGet(ctx context.Context, clt *core.SDKClient, req *idle.AlibabaIdleTenderAftersaleOrderGetAPIRequest, resp *idle.AlibabaIdleTenderAftersaleOrderGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
