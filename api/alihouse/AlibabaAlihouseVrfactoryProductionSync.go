package alihouse

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/alihouse"
)

// AlibabaAlihouseVrfactoryProductionSync vr生产数据上翻
// alibaba.alihouse.vrfactory.production.sync
//
// vr生产数据上翻
func AlibabaAlihouseVrfactoryProductionSync(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseVrfactoryProductionSyncAPIRequest, resp *alihouse.AlibabaAlihouseVrfactoryProductionSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
