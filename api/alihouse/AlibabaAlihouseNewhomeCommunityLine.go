package alihouse

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeCommunityLine 小区上下架
// alibaba.alihouse.newhome.community.line
//
// 小区上下架
func AlibabaAlihouseNewhomeCommunityLine(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeCommunityLineAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeCommunityLineAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
