package tmallcar

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tmallcar"
)

// TmallAliautoServiceReceiptGet isv查询服务工单详情
// tmall.aliauto.service.receipt.get
//
// isv查询服务工单详情
func TmallAliautoServiceReceiptGet(ctx context.Context, clt *core.SDKClient, req *tmallcar.TmallAliautoServiceReceiptGetAPIRequest, resp *tmallcar.TmallAliautoServiceReceiptGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
