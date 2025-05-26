package simba

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/simba"
)

// TaobaoSimbaKeywordsDelete 删除一批关键词
// taobao.simba.keywords.delete
//
// 删除一批关键词
func TaobaoSimbaKeywordsDelete(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaKeywordsDeleteAPIRequest, resp *simba.TaobaoSimbaKeywordsDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
