package xhotelitem

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/xhotelitem"
)

// TaobaoXhotelRoomtypeAdd 房型新增接口（ID重复变更新）
// taobao.xhotel.roomtype.add
//
// 房型添加或更新
func TaobaoXhotelRoomtypeAdd(ctx context.Context, clt *core.SDKClient, req *xhotelitem.TaobaoXhotelRoomtypeAddAPIRequest, resp *xhotelitem.TaobaoXhotelRoomtypeAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
