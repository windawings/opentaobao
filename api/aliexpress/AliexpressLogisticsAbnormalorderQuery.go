package aliexpress

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/aliexpress"
)

// AliexpressLogisticsAbnormalorderQuery 异常订单查询
// aliexpress.logistics.abnormalorder.query
//
// 异常订单查询
func AliexpressLogisticsAbnormalorderQuery(ctx context.Context, clt *core.SDKClient, req *aliexpress.AliexpressLogisticsAbnormalorderQueryAPIRequest, resp *aliexpress.AliexpressLogisticsAbnormalorderQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
