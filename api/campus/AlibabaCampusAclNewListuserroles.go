package campus

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/campus"
)

// AlibabaCampusAclNewListuserroles 查询并标记用户选择的角色
// alibaba.campus.acl.new.listuserroles
//
// 查询并标记用户选择的角色
func AlibabaCampusAclNewListuserroles(ctx context.Context, clt *core.SDKClient, req *campus.AlibabaCampusAclNewListuserrolesAPIRequest, resp *campus.AlibabaCampusAclNewListuserrolesAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
