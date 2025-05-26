package alsc

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/alsc"
)

// AlibabaAlscCrmRuleQuerympricerule 查询品牌下的会员价规则
// alibaba.alsc.crm.rule.querympricerule
//
// 查询品牌下的会员价规则
func AlibabaAlscCrmRuleQuerympricerule(ctx context.Context, clt *core.SDKClient, req *alsc.AlibabaAlscCrmRuleQuerympriceruleAPIRequest, resp *alsc.AlibabaAlscCrmRuleQuerympriceruleAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
