package alihealth2

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/alihealth2"
)

// TaobaoDrugPriceBatchUpdate 商家批量更新宝贝价格
// taobao.drug.price.batch.update
//
// 商家批量更新宝贝价格
func TaobaoDrugPriceBatchUpdate(ctx context.Context, clt *core.SDKClient, req *alihealth2.TaobaoDrugPriceBatchUpdateAPIRequest, resp *alihealth2.TaobaoDrugPriceBatchUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
