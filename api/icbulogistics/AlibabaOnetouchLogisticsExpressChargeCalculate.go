package icbulogistics

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/icbulogistics"
)

// AlibabaOnetouchLogisticsExpressChargeCalculate 计算快递运费&下单参数校验
// alibaba.onetouch.logistics.express.charge.calculate
//
// 计算快递运费、下单参数校验
func AlibabaOnetouchLogisticsExpressChargeCalculate(ctx context.Context, clt *core.SDKClient, req *icbulogistics.AlibabaOnetouchLogisticsExpressChargeCalculateAPIRequest, resp *icbulogistics.AlibabaOnetouchLogisticsExpressChargeCalculateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
