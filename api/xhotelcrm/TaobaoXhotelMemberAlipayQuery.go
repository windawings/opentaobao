package xhotelcrm

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/xhotelcrm"
)

// TaobaoXhotelMemberAlipayQuery 希尔顿会员查询
// taobao.xhotel.member.alipay.query
//
// 希尔顿会员查询
func TaobaoXhotelMemberAlipayQuery(ctx context.Context, clt *core.SDKClient, req *xhotelcrm.TaobaoXhotelMemberAlipayQueryAPIRequest, resp *xhotelcrm.TaobaoXhotelMemberAlipayQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
