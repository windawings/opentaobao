package logistic

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/logistic"
)

// TaobaoLogisticsAddressModify 卖家地址库修改
// taobao.logistics.address.modify
//
// 卖家地址库修改
func TaobaoLogisticsAddressModify(ctx context.Context, clt *core.SDKClient, req *logistic.TaobaoLogisticsAddressModifyAPIRequest, resp *logistic.TaobaoLogisticsAddressModifyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
