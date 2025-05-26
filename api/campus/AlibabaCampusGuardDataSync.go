package campus

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/campus"
)

// AlibabaCampusGuardDataSync 卡巴数据同步
// alibaba.campus.guard.data.sync
//
// 数据同步门禁系统
func AlibabaCampusGuardDataSync(ctx context.Context, clt *core.SDKClient, req *campus.AlibabaCampusGuardDataSyncAPIRequest, resp *campus.AlibabaCampusGuardDataSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
