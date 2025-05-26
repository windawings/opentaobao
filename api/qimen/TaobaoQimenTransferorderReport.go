package qimen

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/qimen"
)

// TaobaoQimenTransferorderReport 调拨单通知
// taobao.qimen.transferorder.report
//
// 调拨单通知
func TaobaoQimenTransferorderReport(ctx context.Context, clt *core.SDKClient, req *qimen.TaobaoQimenTransferorderReportAPIRequest, resp *qimen.TaobaoQimenTransferorderReportAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
