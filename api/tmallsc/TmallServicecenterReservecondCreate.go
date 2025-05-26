package tmallsc

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tmallsc"
)

// TmallServicecenterReservecondCreate 创建主动预约开通条件
// tmall.servicecenter.reservecond.create
//
// 1、设置主动预约开通条件
func TmallServicecenterReservecondCreate(ctx context.Context, clt *core.SDKClient, req *tmallsc.TmallServicecenterReservecondCreateAPIRequest, resp *tmallsc.TmallServicecenterReservecondCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
