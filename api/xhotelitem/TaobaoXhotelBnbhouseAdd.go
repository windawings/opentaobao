package xhotelitem

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/xhotelitem"
)

// TaobaoXhotelBnbhouseAdd 民宿门店信息添加
// taobao.xhotel.bnbhouse.add
//
// 添加和更新民宿门店的信息
func TaobaoXhotelBnbhouseAdd(ctx context.Context, clt *core.SDKClient, req *xhotelitem.TaobaoXhotelBnbhouseAddAPIRequest, resp *xhotelitem.TaobaoXhotelBnbhouseAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
