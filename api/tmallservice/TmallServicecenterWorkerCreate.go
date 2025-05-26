package tmallservice

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tmallservice"
)

// TmallServicecenterWorkerCreate 服务商工人信息创建
// tmall.servicecenter.worker.create
//
// 服务商工人信息创建
func TmallServicecenterWorkerCreate(ctx context.Context, clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkerCreateAPIRequest, resp *tmallservice.TmallServicecenterWorkerCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
