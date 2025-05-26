package drug

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/drug"
)

// TaobaoAlihealthDrugStoreSearch 药品店内搜索
// taobao.alihealth.drug.store.search
//
// 提供给千牛智能客服，在阿里健康O2O店铺内搜索药品
func TaobaoAlihealthDrugStoreSearch(ctx context.Context, clt *core.SDKClient, req *drug.TaobaoAlihealthDrugStoreSearchAPIRequest, resp *drug.TaobaoAlihealthDrugStoreSearchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
