package flightuppc

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/flightuppc"
)

// AlitripFlightExternalAlipaySign 支付宝小程序验签
// alitrip.flight.external.alipay.sign
//
// 支付宝小程序验签
func AlitripFlightExternalAlipaySign(ctx context.Context, clt *core.SDKClient, req *flightuppc.AlitripFlightExternalAlipaySignAPIRequest, resp *flightuppc.AlitripFlightExternalAlipaySignAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
