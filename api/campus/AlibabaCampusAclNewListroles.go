package campus

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/campus"
)

// AlibabaCampusAclNewListroles 查询全部角色
// alibaba.campus.acl.new.listroles
//
// 查询全部角色
func AlibabaCampusAclNewListroles(ctx context.Context, clt *core.SDKClient, req *campus.AlibabaCampusAclNewListrolesAPIRequest, resp *campus.AlibabaCampusAclNewListrolesAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
