package legalcase

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/legalcase"
)

// AlibabaLegalCaseCommonNotice 消息通知
// alibaba.legal.case.common.notice
//
// 同步通知给诉讼系统
func AlibabaLegalCaseCommonNotice(ctx context.Context, clt *core.SDKClient, req *legalcase.AlibabaLegalCaseCommonNoticeAPIRequest, resp *legalcase.AlibabaLegalCaseCommonNoticeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
