package drugtrace

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/drugtrace"
)

// AlibabaAlihealthTracecodesearchGetshowurlVivo 获取药品扫码落地页vivo
// alibaba.alihealth.tracecodesearch.getshowurl.vivo
//
// 获取药品扫码落地页vivo
func AlibabaAlihealthTracecodesearchGetshowurlVivo(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthTracecodesearchGetshowurlVivoAPIRequest, resp *drugtrace.AlibabaAlihealthTracecodesearchGetshowurlVivoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
