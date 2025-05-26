package tmallsc

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tmallsc"
)

// AlibabaMsfserviceWorkerQueryid 查询师傅workerid
// alibaba.msfservice.worker.queryid
//
// 查询师傅workerid
func AlibabaMsfserviceWorkerQueryid(ctx context.Context, clt *core.SDKClient, req *tmallsc.AlibabaMsfserviceWorkerQueryidAPIRequest, resp *tmallsc.AlibabaMsfserviceWorkerQueryidAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
