package simba

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/simba"
)

// TaobaoSubwayCiaGet 查询单元智能出价信息
// taobao.subway.cia.get
//
// 查询单元智能出价信息
func TaobaoSubwayCiaGet(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSubwayCiaGetAPIRequest, resp *simba.TaobaoSubwayCiaGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
