package alsc

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/alsc"
)

// AlibabaAlscCrmCardQueryTemplate 查询卡模板详情
// alibaba.alsc.crm.card.query.template
//
// 查询卡模板详情
func AlibabaAlscCrmCardQueryTemplate(ctx context.Context, clt *core.SDKClient, req *alsc.AlibabaAlscCrmCardQueryTemplateAPIRequest, resp *alsc.AlibabaAlscCrmCardQueryTemplateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
