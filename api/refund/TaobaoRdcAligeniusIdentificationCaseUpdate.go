package refund

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/refund"
)

// TaobaoRdcAligeniusIdentificationCaseUpdate 鉴定工单信息同步
// taobao.rdc.aligenius.identification.case.update
//
// 同步商家鉴定工单信息
func TaobaoRdcAligeniusIdentificationCaseUpdate(ctx context.Context, clt *core.SDKClient, req *refund.TaobaoRdcAligeniusIdentificationCaseUpdateAPIRequest, resp *refund.TaobaoRdcAligeniusIdentificationCaseUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
