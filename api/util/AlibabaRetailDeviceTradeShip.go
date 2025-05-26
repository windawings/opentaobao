package util

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/util"
)

// AlibabaRetailDeviceTradeShip 贩卖机掉货成功通知
// alibaba.retail.device.trade.ship
//
// 贩卖机发货
func AlibabaRetailDeviceTradeShip(ctx context.Context, clt *core.SDKClient, req *util.AlibabaRetailDeviceTradeShipAPIRequest, resp *util.AlibabaRetailDeviceTradeShipAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
