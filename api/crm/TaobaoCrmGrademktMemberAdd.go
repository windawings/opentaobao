package crm

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/crm"
)

// TaobaoCrmGrademktMemberAdd 会员等级营销-会员吸纳
// taobao.crm.grademkt.member.add
//
// 商家通过该接口吸纳线上店铺会员。
func TaobaoCrmGrademktMemberAdd(ctx context.Context, clt *core.SDKClient, req *crm.TaobaoCrmGrademktMemberAddAPIRequest, resp *crm.TaobaoCrmGrademktMemberAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
