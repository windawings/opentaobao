package campus

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/campus"
)

// AlibabaCampusGuardantDataSync 刷卡数据同步
// alibaba.campus.guardant.data.sync
//
// 数据同步门禁系统
func AlibabaCampusGuardantDataSync(ctx context.Context, clt *core.SDKClient, req *campus.AlibabaCampusGuardantDataSyncAPIRequest, resp *campus.AlibabaCampusGuardantDataSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
