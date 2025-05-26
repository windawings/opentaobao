package campus

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/campus"
)

// AlibabaCampusAclNewCheckuserrole 校验用户是否有角色
// alibaba.campus.acl.new.checkuserrole
//
// 校验用户是否有角色
func AlibabaCampusAclNewCheckuserrole(ctx context.Context, clt *core.SDKClient, req *campus.AlibabaCampusAclNewCheckuserroleAPIRequest, resp *campus.AlibabaCampusAclNewCheckuserroleAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
