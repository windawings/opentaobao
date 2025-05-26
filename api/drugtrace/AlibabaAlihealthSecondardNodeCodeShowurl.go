package drugtrace

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/drugtrace"
)

// AlibabaAlihealthSecondardNodeCodeShowurl 查询码信息url
// alibaba.alihealth.secondard.node.code.showurl
//
// 二级节点查询码信息url
func AlibabaAlihealthSecondardNodeCodeShowurl(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthSecondardNodeCodeShowurlAPIRequest, resp *drugtrace.AlibabaAlihealthSecondardNodeCodeShowurlAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
