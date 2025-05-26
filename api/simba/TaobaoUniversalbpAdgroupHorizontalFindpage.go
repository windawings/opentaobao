package simba

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/simba"
)

// TaobaoUniversalbpAdgroupHorizontalFindpage 查询单元分页列表
// taobao.universalbp.adgroup.horizontal.findpage
//
// 查询单元分页列表
func TaobaoUniversalbpAdgroupHorizontalFindpage(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoUniversalbpAdgroupHorizontalFindpageAPIRequest, resp *simba.TaobaoUniversalbpAdgroupHorizontalFindpageAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
