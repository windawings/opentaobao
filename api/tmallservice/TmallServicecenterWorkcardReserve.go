package tmallservice

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tmallservice"
)

// TmallServicecenterWorkcardReserve 工单预约
// tmall.servicecenter.workcard.reserve
//
// 服务工单更新通用接口
func TmallServicecenterWorkcardReserve(ctx context.Context, clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkcardReserveAPIRequest, resp *tmallservice.TmallServicecenterWorkcardReserveAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
