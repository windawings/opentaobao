package eticket

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/eticket"
)

// TaobaoEticketMerchantMaSend 码商发码成功回调接口
// taobao.eticket.merchant.ma.send
//
// 码商发码成功回调接口
func TaobaoEticketMerchantMaSend(ctx context.Context, clt *core.SDKClient, req *eticket.TaobaoEticketMerchantMaSendAPIRequest, resp *eticket.TaobaoEticketMerchantMaSendAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
