package campus

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/campus"
)

// AlibabaCampusAclGetmenubyempid 查询用户的菜单
// alibaba.campus.acl.getmenubyempid
//
// 查询用户的菜单
func AlibabaCampusAclGetmenubyempid(ctx context.Context, clt *core.SDKClient, req *campus.AlibabaCampusAclGetmenubyempidAPIRequest, resp *campus.AlibabaCampusAclGetmenubyempidAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
