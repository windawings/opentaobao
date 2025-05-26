package alihouse

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/alihouse"
)

// AlibabaAlihouseProjectActivitySync 电商券数据同步
// alibaba.alihouse.project.activity.sync
//
// 电商券数据同步
func AlibabaAlihouseProjectActivitySync(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseProjectActivitySyncAPIRequest, resp *alihouse.AlibabaAlihouseProjectActivitySyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
