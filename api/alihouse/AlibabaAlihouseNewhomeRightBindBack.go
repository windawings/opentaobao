package alihouse

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeRightBindBack 权限回流
// alibaba.alihouse.newhome.right.bind.back
//
// 权限回流
func AlibabaAlihouseNewhomeRightBindBack(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeRightBindBackAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeRightBindBackAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
