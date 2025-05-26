package simba

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/simba"
)

// TaobaoSimbaAdgroupDelete 删除一个推广组
// taobao.simba.adgroup.delete
//
// 删除一个推广组
func TaobaoSimbaAdgroupDelete(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaAdgroupDeleteAPIRequest, resp *simba.TaobaoSimbaAdgroupDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
