package alihouse

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeProjectSalestime 楼盘销售时刻同步
// alibaba.alihouse.newhome.project.salestime
//
// 楼盘销售时刻同步
func AlibabaAlihouseNewhomeProjectSalestime(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeProjectSalestimeAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeProjectSalestimeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
