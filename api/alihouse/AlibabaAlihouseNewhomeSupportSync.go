package alihouse

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeSupportSync 周边配套数据同步
// alibaba.alihouse.newhome.support.sync
//
// 周边配套数据同步
func AlibabaAlihouseNewhomeSupportSync(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeSupportSyncAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeSupportSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
