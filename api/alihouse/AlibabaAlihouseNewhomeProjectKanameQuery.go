package alihouse

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeProjectKanameQuery 查询KA楼盘名称
// alibaba.alihouse.newhome.project.kaname.query
//
// 查询KA楼盘名称
func AlibabaAlihouseNewhomeProjectKanameQuery(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeProjectKanameQueryAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeProjectKanameQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
