package simba

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/simba"
)

// TaobaoUniversalbpCreativeHorizontalFindpage 横向管理创意分页查询
// taobao.universalbp.creative.horizontal.findpage
//
// 横向管理创意分页查询
func TaobaoUniversalbpCreativeHorizontalFindpage(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoUniversalbpCreativeHorizontalFindpageAPIRequest, resp *simba.TaobaoUniversalbpCreativeHorizontalFindpageAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
