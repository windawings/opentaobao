package tmallservice

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tmallservice"
)

// TmallServicecenterWorkcardReservefail 预约失败
// tmall.servicecenter.workcard.reservefail
//
// 服务商调用该接口回传工单预约失败
func TmallServicecenterWorkcardReservefail(ctx context.Context, clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkcardReservefailAPIRequest, resp *tmallservice.TmallServicecenterWorkcardReservefailAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
