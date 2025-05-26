package mtopopen

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/mtopopen"
)

// TaobaoLogisticsAppletPackageQuery 淘宝包裹查询
// taobao.logistics.applet.package.query
//
// 淘宝包裹查询
func TaobaoLogisticsAppletPackageQuery(ctx context.Context, clt *core.SDKClient, req *mtopopen.TaobaoLogisticsAppletPackageQueryAPIRequest, resp *mtopopen.TaobaoLogisticsAppletPackageQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
