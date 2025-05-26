package simba

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/simba"
)

// TaobaoSimbaSalestarAdgroupFindbycampid (销量明星)批量获取推广计划下的推广组信息
// taobao.simba.salestar.adgroup.findbycampid
//
// 批量得到推广计划下的推广组
func TaobaoSimbaSalestarAdgroupFindbycampid(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaSalestarAdgroupFindbycampidAPIRequest, resp *simba.TaobaoSimbaSalestarAdgroupFindbycampidAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
