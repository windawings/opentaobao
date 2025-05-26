package wlbimports

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/wlbimports"
)

// CainiaoGlobalImPickupAppointmentOrderStatus 预约单状态查询
// cainiao.global.im.pickup.appointment.order.status
//
// 预约单状态查询
func CainiaoGlobalImPickupAppointmentOrderStatus(ctx context.Context, clt *core.SDKClient, req *wlbimports.CainiaoGlobalImPickupAppointmentOrderStatusAPIRequest, resp *wlbimports.CainiaoGlobalImPickupAppointmentOrderStatusAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
