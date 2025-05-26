package icbulogistics

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/icbulogistics"
)

// AlibabaOnetouchLogisticsExpressAddressProvinceList 四级地址库-省
// alibaba.onetouch.logistics.express.address.province.list
//
// 四级地址库-省
func AlibabaOnetouchLogisticsExpressAddressProvinceList(ctx context.Context, clt *core.SDKClient, req *icbulogistics.AlibabaOnetouchLogisticsExpressAddressProvinceListAPIRequest, resp *icbulogistics.AlibabaOnetouchLogisticsExpressAddressProvinceListAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
