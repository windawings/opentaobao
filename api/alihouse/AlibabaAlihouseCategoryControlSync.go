package alihouse

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/alihouse"
)

// AlibabaAlihouseCategoryControlSync 类目权限上翻
// alibaba.alihouse.category.control.sync
//
// 类目权限上翻
func AlibabaAlihouseCategoryControlSync(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseCategoryControlSyncAPIRequest, resp *alihouse.AlibabaAlihouseCategoryControlSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
