package campus

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/campus"
)

// AlibabaCampusAclNewGetrolewithmenutreenodes 根据角色id查询权限
// alibaba.campus.acl.new.getrolewithmenutreenodes
//
// 根据角色id查询权限
func AlibabaCampusAclNewGetrolewithmenutreenodes(ctx context.Context, clt *core.SDKClient, req *campus.AlibabaCampusAclNewGetrolewithmenutreenodesAPIRequest, resp *campus.AlibabaCampusAclNewGetrolewithmenutreenodesAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
