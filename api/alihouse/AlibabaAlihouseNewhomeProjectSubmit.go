package alihouse

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeProjectSubmit 提交楼盘信息
// alibaba.alihouse.newhome.project.submit
//
// 提交楼盘信息
func AlibabaAlihouseNewhomeProjectSubmit(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeProjectSubmitAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeProjectSubmitAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
