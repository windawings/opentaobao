package tmallhk

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tmallhk"
)

// TmallTraceplatformTicketOrderUpload 上传小票数据
// tmall.traceplatform.ticket.order.upload
//
// upsertOrderBySeller
func TmallTraceplatformTicketOrderUpload(ctx context.Context, clt *core.SDKClient, req *tmallhk.TmallTraceplatformTicketOrderUploadAPIRequest, resp *tmallhk.TmallTraceplatformTicketOrderUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
