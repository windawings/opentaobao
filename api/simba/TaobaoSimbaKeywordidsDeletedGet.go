package simba

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/simba"
)

// TaobaoSimbaKeywordidsDeletedGet 获取删除的词ID
// taobao.simba.keywordids.deleted.get
//
// 获取删除的词ID
func TaobaoSimbaKeywordidsDeletedGet(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaKeywordidsDeletedGetAPIRequest, resp *simba.TaobaoSimbaKeywordidsDeletedGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
