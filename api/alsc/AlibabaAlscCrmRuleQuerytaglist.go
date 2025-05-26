package alsc

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/alsc"
)

// AlibabaAlscCrmRuleQuerytaglist 查询标签列表
// alibaba.alsc.crm.rule.querytaglist
//
// 查询标签列表
func AlibabaAlscCrmRuleQuerytaglist(ctx context.Context, clt *core.SDKClient, req *alsc.AlibabaAlscCrmRuleQuerytaglistAPIRequest, resp *alsc.AlibabaAlscCrmRuleQuerytaglistAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
