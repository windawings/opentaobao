package alihouse

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeCasefieldActivitySubmit 案场活动维护
// alibaba.alihouse.newhome.casefield.activity.submit
//
// 案场活动维护
func AlibabaAlihouseNewhomeCasefieldActivitySubmit(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeCasefieldActivitySubmitAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeCasefieldActivitySubmitAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
