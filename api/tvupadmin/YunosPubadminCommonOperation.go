package tvupadmin

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tvupadmin"
)

// YunosPubadminCommonOperation 内部迎客松通用服务
// yunos.pubadmin.common.operation
//
// 内部迎客松通用服务
func YunosPubadminCommonOperation(ctx context.Context, clt *core.SDKClient, req *tvupadmin.YunosPubadminCommonOperationAPIRequest, resp *tvupadmin.YunosPubadminCommonOperationAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
