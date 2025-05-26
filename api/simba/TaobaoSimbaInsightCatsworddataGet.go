package simba

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/simba"
)

// TaobaoSimbaInsightCatsworddataGet 获取类目下关键词的数据
// taobao.simba.insight.catsworddata.get
//
// 获取给定词在给定类目下的详细数据
func TaobaoSimbaInsightCatsworddataGet(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaInsightCatsworddataGetAPIRequest, resp *simba.TaobaoSimbaInsightCatsworddataGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
