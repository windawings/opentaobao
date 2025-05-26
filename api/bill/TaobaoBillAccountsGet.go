package bill

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/bill"
)

// TaobaoBillAccountsGet 查询费用科目信息(限自研商家)
// taobao.bill.accounts.get
//
// 查询费用账户信息
func TaobaoBillAccountsGet(ctx context.Context, clt *core.SDKClient, req *bill.TaobaoBillAccountsGetAPIRequest, resp *bill.TaobaoBillAccountsGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
