package maitix

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/maitix"
)

// AlibabaDamaiMaitixOrderCancel 大麦-库存释放
// alibaba.damai.maitix.order.cancel
//
// 库存释放
func AlibabaDamaiMaitixOrderCancel(ctx context.Context, clt *core.SDKClient, req *maitix.AlibabaDamaiMaitixOrderCancelAPIRequest, resp *maitix.AlibabaDamaiMaitixOrderCancelAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
