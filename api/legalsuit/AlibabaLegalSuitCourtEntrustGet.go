package legalsuit

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/legalsuit"
)

// AlibabaLegalSuitCourtEntrustGet 委托开庭服务查询
// alibaba.legal.suit.court.entrust.get
//
// 查询委托开庭信息
func AlibabaLegalSuitCourtEntrustGet(ctx context.Context, clt *core.SDKClient, req *legalsuit.AlibabaLegalSuitCourtEntrustGetAPIRequest, resp *legalsuit.AlibabaLegalSuitCourtEntrustGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
