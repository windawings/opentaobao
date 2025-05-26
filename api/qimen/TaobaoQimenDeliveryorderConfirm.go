package qimen

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/qimen"
)

// TaobaoQimenDeliveryorderConfirm 发货单确认接口
// taobao.qimen.deliveryorder.confirm
//
// taobao.qimen.deliveryorder.confirm
func TaobaoQimenDeliveryorderConfirm(ctx context.Context, clt *core.SDKClient, req *qimen.TaobaoQimenDeliveryorderConfirmAPIRequest, resp *qimen.TaobaoQimenDeliveryorderConfirmAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
