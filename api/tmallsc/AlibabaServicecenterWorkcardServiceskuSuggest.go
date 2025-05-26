package tmallsc

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tmallsc"
)

// AlibabaServicecenterWorkcardServiceskuSuggest 服务商反馈需要履行的服务项
// alibaba.servicecenter.workcard.servicesku.suggest
//
// 服务商反馈需要履行的服务项
func AlibabaServicecenterWorkcardServiceskuSuggest(ctx context.Context, clt *core.SDKClient, req *tmallsc.AlibabaServicecenterWorkcardServiceskuSuggestAPIRequest, resp *tmallsc.AlibabaServicecenterWorkcardServiceskuSuggestAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
