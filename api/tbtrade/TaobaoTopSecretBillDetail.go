package tbtrade

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tbtrade"
)

// TaobaoTopSecretBillDetail 服务商的商家解密账单详情查询
// taobao.top.secret.bill.detail
//
// 服务商的商家解密账单详情查询，仅对90天内的账单提供SLA保障。
func TaobaoTopSecretBillDetail(ctx context.Context, clt *core.SDKClient, req *tbtrade.TaobaoTopSecretBillDetailAPIRequest, resp *tbtrade.TaobaoTopSecretBillDetailAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
