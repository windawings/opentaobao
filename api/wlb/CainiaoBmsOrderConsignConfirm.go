package wlb

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/wlb"
)

// CainiaoBmsOrderConsignConfirm BMS出库通知
// cainiao.bms.order.consign.confirm
//
// BMS出库后，通知ISV
func CainiaoBmsOrderConsignConfirm(ctx context.Context, clt *core.SDKClient, req *wlb.CainiaoBmsOrderConsignConfirmAPIRequest, resp *wlb.CainiaoBmsOrderConsignConfirmAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
