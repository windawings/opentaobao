package omniorder

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/omniorder"
)

// TaobaoOmniorderStoreCollectconfigGet 查询门店自提配置内容
// taobao.omniorder.store.collectconfig.get
//
// 查询门店自提配置内容
func TaobaoOmniorderStoreCollectconfigGet(ctx context.Context, clt *core.SDKClient, req *omniorder.TaobaoOmniorderStoreCollectconfigGetAPIRequest, resp *omniorder.TaobaoOmniorderStoreCollectconfigGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
