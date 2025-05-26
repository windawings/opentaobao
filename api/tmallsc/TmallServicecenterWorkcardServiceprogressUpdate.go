package tmallsc

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tmallsc"
)

// TmallServicecenterWorkcardServiceprogressUpdate 回传工单服务进度
// tmall.servicecenter.workcard.serviceprogress.update
//
// 回传工单服务进度
func TmallServicecenterWorkcardServiceprogressUpdate(ctx context.Context, clt *core.SDKClient, req *tmallsc.TmallServicecenterWorkcardServiceprogressUpdateAPIRequest, resp *tmallsc.TmallServicecenterWorkcardServiceprogressUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
