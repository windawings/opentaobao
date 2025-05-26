package alihouse

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeHouseCooperateBrandUpdate 租房合作品牌更新接口
// alibaba.alihouse.existinghome.house.cooperate.brand.update
//
// 租房合作品牌更新接口
func AlibabaAlihouseExistinghomeHouseCooperateBrandUpdate(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeHouseCooperateBrandUpdateAPIRequest, resp *alihouse.AlibabaAlihouseExistinghomeHouseCooperateBrandUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
