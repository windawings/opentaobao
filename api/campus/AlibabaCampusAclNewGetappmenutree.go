package campus

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/campus"
)

// AlibabaCampusAclNewGetappmenutree 查询应用下的菜单树
// alibaba.campus.acl.new.getappmenutree
//
// 查询应用下的菜单树
func AlibabaCampusAclNewGetappmenutree(ctx context.Context, clt *core.SDKClient, req *campus.AlibabaCampusAclNewGetappmenutreeAPIRequest, resp *campus.AlibabaCampusAclNewGetappmenutreeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
