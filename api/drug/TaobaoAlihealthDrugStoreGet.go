package drug

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/drug"
)

// TaobaoAlihealthDrugStoreGet 根据店铺id获取店铺详情
// taobao.alihealth.drug.store.get
//
// 根据店铺id获取店铺详情
func TaobaoAlihealthDrugStoreGet(ctx context.Context, clt *core.SDKClient, req *drug.TaobaoAlihealthDrugStoreGetAPIRequest, resp *drug.TaobaoAlihealthDrugStoreGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
