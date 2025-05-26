package legalcase

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/legalcase"
)

// AlibabaLegalCaseEntrustGet 委托
// alibaba.legal.case.entrust.get
//
// 获取委托案件的基本信息
func AlibabaLegalCaseEntrustGet(ctx context.Context, clt *core.SDKClient, req *legalcase.AlibabaLegalCaseEntrustGetAPIRequest, resp *legalcase.AlibabaLegalCaseEntrustGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
