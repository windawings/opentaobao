package fenxiao

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/fenxiao"
)

// AlibabaAscpCnskuSearch 供应链中台货品搜索接口
// alibaba.ascp.cnsku.search
//
// 供应链中台货品搜索接口
func AlibabaAscpCnskuSearch(ctx context.Context, clt *core.SDKClient, req *fenxiao.AlibabaAscpCnskuSearchAPIRequest, resp *fenxiao.AlibabaAscpCnskuSearchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
