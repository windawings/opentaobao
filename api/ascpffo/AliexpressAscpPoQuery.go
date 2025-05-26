package ascpffo

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/ascpffo"
)

// AliexpressAscpPoQuery AliExpress采购单查询API
// aliexpress.ascp.po.query
//
// AE仓发业务采购单查询
func AliexpressAscpPoQuery(ctx context.Context, clt *core.SDKClient, req *ascpffo.AliexpressAscpPoQueryAPIRequest, resp *ascpffo.AliexpressAscpPoQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
