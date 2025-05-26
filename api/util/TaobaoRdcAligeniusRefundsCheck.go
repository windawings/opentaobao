package util

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/util"
)

// TaobaoRdcAligeniusRefundsCheck 退款信息审核
// taobao.rdc.aligenius.refunds.check
//
// 根据退款信息，对退款单进行审核
func TaobaoRdcAligeniusRefundsCheck(ctx context.Context, clt *core.SDKClient, req *util.TaobaoRdcAligeniusRefundsCheckAPIRequest, resp *util.TaobaoRdcAligeniusRefundsCheckAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
