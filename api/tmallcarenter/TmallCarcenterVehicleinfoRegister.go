package tmallcarenter

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tmallcarenter"
)

// TmallCarcenterVehicleinfoRegister 车型数据更新
// tmall.carcenter.vehicleinfo.register
//
// 基本车型信息维护
func TmallCarcenterVehicleinfoRegister(ctx context.Context, clt *core.SDKClient, req *tmallcarenter.TmallCarcenterVehicleinfoRegisterAPIRequest, resp *tmallcarenter.TmallCarcenterVehicleinfoRegisterAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
