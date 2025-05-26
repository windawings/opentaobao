package wdk

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/wdk"
)

// AlibabaWdkBmPaiyangStatDataQuery 派样统计数据查询
// alibaba.wdk.bm.paiyang.stat.data.query
//
// 派样统计数据查询
func AlibabaWdkBmPaiyangStatDataQuery(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkBmPaiyangStatDataQueryAPIRequest, resp *wdk.AlibabaWdkBmPaiyangStatDataQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
