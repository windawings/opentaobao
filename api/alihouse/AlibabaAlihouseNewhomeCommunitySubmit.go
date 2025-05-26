package alihouse

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeCommunitySubmit 提交小区信息
// alibaba.alihouse.newhome.community.submit
//
// 提交小区信息
func AlibabaAlihouseNewhomeCommunitySubmit(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeCommunitySubmitAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeCommunitySubmitAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
