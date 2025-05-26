package alihouse

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeApartmentOuterid 公寓更新outerid
// alibaba.alihouse.newhome.apartment.outerid
//
// 公寓更新outerid
func AlibabaAlihouseNewhomeApartmentOuterid(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeApartmentOuteridAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeApartmentOuteridAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
