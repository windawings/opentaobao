package campus

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/campus"
)

// AlibabaCampusAclNewCheckuserpermission 校验用户是否有权限
// alibaba.campus.acl.new.checkuserpermission
//
// 校验用户是否有权限
func AlibabaCampusAclNewCheckuserpermission(ctx context.Context, clt *core.SDKClient, req *campus.AlibabaCampusAclNewCheckuserpermissionAPIRequest, resp *campus.AlibabaCampusAclNewCheckuserpermissionAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
