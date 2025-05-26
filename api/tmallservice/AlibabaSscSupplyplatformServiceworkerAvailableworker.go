package tmallservice

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tmallservice"
)

// AlibabaSscSupplyplatformServiceworkerAvailableworker 查询可用工人
// alibaba.ssc.supplyplatform.serviceworker.availableworker
//
// 可用工人查询
func AlibabaSscSupplyplatformServiceworkerAvailableworker(ctx context.Context, clt *core.SDKClient, req *tmallservice.AlibabaSscSupplyplatformServiceworkerAvailableworkerAPIRequest, resp *tmallservice.AlibabaSscSupplyplatformServiceworkerAvailableworkerAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
