package bill

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/bill"
)

// TaobaoTaeAccountsGet tae查询费用科目信息
// taobao.tae.accounts.get
//
// tae查询费用科目信息
func TaobaoTaeAccountsGet(ctx context.Context, clt *core.SDKClient, req *bill.TaobaoTaeAccountsGetAPIRequest, resp *bill.TaobaoTaeAccountsGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
