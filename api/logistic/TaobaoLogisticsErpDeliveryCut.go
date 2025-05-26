package logistic

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/logistic"
)

// TaobaoLogisticsErpDeliveryCut ERP发起配拦截
// taobao.logistics.erp.delivery.cut
//
// ERP发起配拦截
func TaobaoLogisticsErpDeliveryCut(ctx context.Context, clt *core.SDKClient, req *logistic.TaobaoLogisticsErpDeliveryCutAPIRequest, resp *logistic.TaobaoLogisticsErpDeliveryCutAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
