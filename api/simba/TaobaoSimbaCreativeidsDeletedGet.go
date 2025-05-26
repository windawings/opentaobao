package simba

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/simba"
)

// TaobaoSimbaCreativeidsDeletedGet 获取删除的创意ID
// taobao.simba.creativeids.deleted.get
//
// 获取删除的创意ID
func TaobaoSimbaCreativeidsDeletedGet(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaCreativeidsDeletedGetAPIRequest, resp *simba.TaobaoSimbaCreativeidsDeletedGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
