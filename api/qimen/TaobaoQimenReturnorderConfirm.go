package qimen

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/qimen"
)

// TaobaoQimenReturnorderConfirm 退货入库单确认接口
// taobao.qimen.returnorder.confirm
//
// taobao.qimen.returnorder.confirm
func TaobaoQimenReturnorderConfirm(ctx context.Context, clt *core.SDKClient, req *qimen.TaobaoQimenReturnorderConfirmAPIRequest, resp *qimen.TaobaoQimenReturnorderConfirmAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
