package simba

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/simba"
)

// TaobaoSimbaAdgroupOnlineitemsvonGet 获取用户上架在线销售的全部宝贝
// taobao.simba.adgroup.onlineitemsvon.get
//
// 获取用户上架在线销售的全部宝贝
func TaobaoSimbaAdgroupOnlineitemsvonGet(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaAdgroupOnlineitemsvonGetAPIRequest, resp *simba.TaobaoSimbaAdgroupOnlineitemsvonGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
