package mozi

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/mozi"
)

// AlibabaMoziFusionAddorupdateEmployeeAccount 添加人员和账号复合接口
// alibaba.mozi.fusion.addorupdate.employee.account
//
// 添加人员和账号复合接口
func AlibabaMoziFusionAddorupdateEmployeeAccount(ctx context.Context, clt *core.SDKClient, req *mozi.AlibabaMoziFusionAddorupdateEmployeeAccountAPIRequest, resp *mozi.AlibabaMoziFusionAddorupdateEmployeeAccountAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
