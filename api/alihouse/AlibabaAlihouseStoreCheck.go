package alihouse

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/alihouse"
)

// AlibabaAlihouseStoreCheck 门店对账查询工具
// alibaba.alihouse.store.check
//
// 门店对账查询工具
func AlibabaAlihouseStoreCheck(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseStoreCheckAPIRequest, resp *alihouse.AlibabaAlihouseStoreCheckAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
