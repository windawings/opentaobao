package xhotelitem

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/xhotelitem"
)

// TaobaoXhotelBnbpromoAdd 自促活动申请接口
// taobao.xhotel.bnbpromo.add
//
// 自促活动申请接口
func TaobaoXhotelBnbpromoAdd(ctx context.Context, clt *core.SDKClient, req *xhotelitem.TaobaoXhotelBnbpromoAddAPIRequest, resp *xhotelitem.TaobaoXhotelBnbpromoAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
