package xhotelitem

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/xhotelitem"
)

// TaobaoXhotelItemSelectionSellerStatHotshid 供应链选品热销标准酒店覆盖情况
// taobao.xhotel.item.selection.seller.stat.hotshid
//
// 供应链选品热销标准酒店覆盖情况
func TaobaoXhotelItemSelectionSellerStatHotshid(ctx context.Context, clt *core.SDKClient, req *xhotelitem.TaobaoXhotelItemSelectionSellerStatHotshidAPIRequest, resp *xhotelitem.TaobaoXhotelItemSelectionSellerStatHotshidAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
