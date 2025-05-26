package campus

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/campus"
)

// AlibabaCampusGuardControllerOfflinelog 门禁控制器离线日志同步
// alibaba.campus.guard.controller.offlinelog
//
// 门禁控制器离线日志同步
func AlibabaCampusGuardControllerOfflinelog(ctx context.Context, clt *core.SDKClient, req *campus.AlibabaCampusGuardControllerOfflinelogAPIRequest, resp *campus.AlibabaCampusGuardControllerOfflinelogAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
