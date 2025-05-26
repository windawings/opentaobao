package alihouse

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/alihouse"
)

// AlibabaAlihouseAdminThemeCreate 创建云主题
// alibaba.alihouse.admin.theme.create
//
// 创建云主题
func AlibabaAlihouseAdminThemeCreate(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseAdminThemeCreateAPIRequest, resp *alihouse.AlibabaAlihouseAdminThemeCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
