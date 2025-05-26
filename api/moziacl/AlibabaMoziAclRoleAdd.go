package moziacl

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/moziacl"
)

// AlibabaMoziAclRoleAdd 新增一个角色
// alibaba.mozi.acl.role.add
//
// 新增一个角色
func AlibabaMoziAclRoleAdd(ctx context.Context, clt *core.SDKClient, req *moziacl.AlibabaMoziAclRoleAddAPIRequest, resp *moziacl.AlibabaMoziAclRoleAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
