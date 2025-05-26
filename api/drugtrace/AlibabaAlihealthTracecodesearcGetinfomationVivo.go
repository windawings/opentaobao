package drugtrace

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/drugtrace"
)

// AlibabaAlihealthTracecodesearcGetinfomationVivo 获取vivo banner
// alibaba.alihealth.tracecodesearc.getinfomation.vivo
//
// 获取vivo banner  url
func AlibabaAlihealthTracecodesearcGetinfomationVivo(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthTracecodesearcGetinfomationVivoAPIRequest, resp *drugtrace.AlibabaAlihealthTracecodesearcGetinfomationVivoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
