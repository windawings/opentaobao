package tmallservice

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tmallservice"
)

// AlibabaSscSupplyplatformServiceworkerQuitservicestore 工人退出网点
// alibaba.ssc.supplyplatform.serviceworker.quitservicestore
//
// 工人退出网点
func AlibabaSscSupplyplatformServiceworkerQuitservicestore(ctx context.Context, clt *core.SDKClient, req *tmallservice.AlibabaSscSupplyplatformServiceworkerQuitservicestoreAPIRequest, resp *tmallservice.AlibabaSscSupplyplatformServiceworkerQuitservicestoreAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
