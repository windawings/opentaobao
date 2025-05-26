package alihouse

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeProjectPresalepermitDelete 删除楼盘预售证
// alibaba.alihouse.newhome.project.presalepermit.delete
//
// 删除楼盘预售证信息
func AlibabaAlihouseNewhomeProjectPresalepermitDelete(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeProjectPresalepermitDeleteAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeProjectPresalepermitDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
