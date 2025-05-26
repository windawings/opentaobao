package happytrip

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/happytrip"
)

// AlibabaHappytripTaxiOrderCancel 取消叫车
// alibaba.happytrip.taxi.order.cancel
//
// 取消叫车订单,行程中的订单不能取消
func AlibabaHappytripTaxiOrderCancel(ctx context.Context, clt *core.SDKClient, req *happytrip.AlibabaHappytripTaxiOrderCancelAPIRequest, resp *happytrip.AlibabaHappytripTaxiOrderCancelAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
