package alsc

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/alsc"
)

// AlibabaAlscCrmRuleQueryoptplan 查询运营计划
// alibaba.alsc.crm.rule.queryoptplan
//
// 查询运营计划
func AlibabaAlscCrmRuleQueryoptplan(ctx context.Context, clt *core.SDKClient, req *alsc.AlibabaAlscCrmRuleQueryoptplanAPIRequest, resp *alsc.AlibabaAlscCrmRuleQueryoptplanAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
