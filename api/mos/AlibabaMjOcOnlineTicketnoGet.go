package mos

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/mos"
)

// AlibabaMjOcOnlineTicketnoGet 线上小票号获取
// alibaba.mj.oc.online.ticketno.get
//
// 线上小票号获取
func AlibabaMjOcOnlineTicketnoGet(ctx context.Context, clt *core.SDKClient, req *mos.AlibabaMjOcOnlineTicketnoGetAPIRequest, resp *mos.AlibabaMjOcOnlineTicketnoGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
