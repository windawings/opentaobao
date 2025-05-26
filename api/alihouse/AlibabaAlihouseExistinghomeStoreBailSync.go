package alihouse

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeStoreBailSync 门店保证金余额同步
// alibaba.alihouse.existinghome.store.bail.sync
//
// 门店保证金余额同步
func AlibabaAlihouseExistinghomeStoreBailSync(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeStoreBailSyncAPIRequest, resp *alihouse.AlibabaAlihouseExistinghomeStoreBailSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
