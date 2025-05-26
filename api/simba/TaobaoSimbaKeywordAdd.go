package simba

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/simba"
)

// TaobaoSimbaKeywordAdd （新）关键词新增接口
// taobao.simba.keyword.add
//
// （新）关键词更新相关接口
func TaobaoSimbaKeywordAdd(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaKeywordAddAPIRequest, resp *simba.TaobaoSimbaKeywordAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
