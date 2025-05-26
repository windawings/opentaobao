package logistic

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/logistic"
)

// AliexpressLocalLogisticsReportShipped report as shipped
// aliexpress.local.logistics.report.shipped
//
// report as shipped
func AliexpressLocalLogisticsReportShipped(ctx context.Context, clt *core.SDKClient, req *logistic.AliexpressLocalLogisticsReportShippedAPIRequest, resp *logistic.AliexpressLocalLogisticsReportShippedAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
