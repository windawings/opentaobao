package alihouse

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeLayoutSync 房通户型数据同步
// alibaba.alihouse.newhome.layout.sync
//
// 房通户型数据同步
func AlibabaAlihouseNewhomeLayoutSync(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeLayoutSyncAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeLayoutSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
