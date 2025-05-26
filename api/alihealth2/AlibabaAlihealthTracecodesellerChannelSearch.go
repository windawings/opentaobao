package alihealth2

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/alihealth2"
)

// AlibabaAlihealthTracecodesellerChannelSearch 查询渠道商api
// alibaba.alihealth.tracecodeseller.channel.search
//
// 查询渠道商api
func AlibabaAlihealthTracecodesellerChannelSearch(ctx context.Context, clt *core.SDKClient, req *alihealth2.AlibabaAlihealthTracecodesellerChannelSearchAPIRequest, resp *alihealth2.AlibabaAlihealthTracecodesellerChannelSearchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
