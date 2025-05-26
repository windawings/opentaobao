package logistic

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/logistic"
)

// CainiaoMemberCourierCpresign cp清理离职用户信息
// cainiao.member.courier.cpresign
//
// CP清理内部离职的用户信息
func CainiaoMemberCourierCpresign(ctx context.Context, clt *core.SDKClient, req *logistic.CainiaoMemberCourierCpresignAPIRequest, resp *logistic.CainiaoMemberCourierCpresignAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
