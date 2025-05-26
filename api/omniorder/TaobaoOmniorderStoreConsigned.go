package omniorder

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/omniorder"
)

// TaobaoOmniorderStoreConsigned Pos端门店发货
// taobao.omniorder.store.consigned
//
// ISV Pos端门店发货，通知星盘
func TaobaoOmniorderStoreConsigned(ctx context.Context, clt *core.SDKClient, req *omniorder.TaobaoOmniorderStoreConsignedAPIRequest, resp *omniorder.TaobaoOmniorderStoreConsignedAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
