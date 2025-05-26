package xhotelitem

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/xhotelitem"
)

// TaobaoXhotelQuotaUpdate 库存更新接口
// taobao.xhotel.quota.update
//
// 库存更新接口
func TaobaoXhotelQuotaUpdate(ctx context.Context, clt *core.SDKClient, req *xhotelitem.TaobaoXhotelQuotaUpdateAPIRequest, resp *xhotelitem.TaobaoXhotelQuotaUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
