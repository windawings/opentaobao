package ascpchannel

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/ascpchannel"
)

// TaobaoInvTurnoverQuery 业务库存流水查询
// taobao.inv.turnover.query
//
// 业务库存流水
func TaobaoInvTurnoverQuery(ctx context.Context, clt *core.SDKClient, req *ascpchannel.TaobaoInvTurnoverQueryAPIRequest, resp *ascpchannel.TaobaoInvTurnoverQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
