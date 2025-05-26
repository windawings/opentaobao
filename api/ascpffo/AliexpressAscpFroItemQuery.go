package ascpffo

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/ascpffo"
)

// AliexpressAscpFroItemQuery AliExpress销退单明细查询API
// aliexpress.ascp.fro.item.query
//
// AE履约销退单明细查询API
func AliexpressAscpFroItemQuery(ctx context.Context, clt *core.SDKClient, req *ascpffo.AliexpressAscpFroItemQueryAPIRequest, resp *ascpffo.AliexpressAscpFroItemQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
