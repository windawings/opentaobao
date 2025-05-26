package xhotelitem

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/xhotelitem"
)

// TaobaoXhotelHouseAdd 非标准民宿房源添加
// taobao.xhotel.house.add
//
// 添加酒店或更新酒店
func TaobaoXhotelHouseAdd(ctx context.Context, clt *core.SDKClient, req *xhotelitem.TaobaoXhotelHouseAddAPIRequest, resp *xhotelitem.TaobaoXhotelHouseAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
